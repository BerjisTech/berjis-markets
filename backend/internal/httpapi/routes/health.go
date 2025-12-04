package routes

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type dbPinger interface {
	Ping(context.Context) error
}

type cachePinger interface {
	Ping(context.Context) *redis.StatusCmd
}

func RegisterHealth(router *gin.Engine, db dbPinger, cache cachePinger) {
	router.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":    "ok",
			"timestamp": time.Now().UTC().Format(time.RFC3339Nano),
		})
	})

	router.GET("/readyz", func(c *gin.Context) {
		ctx := c.Request.Context()
		dbErr := db.Ping(ctx)
		cacheErr := cache.Ping(ctx).Err()

		if dbErr != nil || cacheErr != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"status":      "degraded",
				"timestamp":   time.Now().UTC().Format(time.RFC3339Nano),
				"db_error":    errString(dbErr),
				"cache_error": errString(cacheErr),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":    "ready",
			"timestamp": time.Now().UTC().Format(time.RFC3339Nano),
		})
	})
}

func errString(err error) string {
	if err == nil {
		return ""
	}
	return err.Error()
}
