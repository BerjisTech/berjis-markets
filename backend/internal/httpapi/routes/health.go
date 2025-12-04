package routes

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func RegisterHealth(group *gin.RouterGroup) {
	group.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":    "ok",
			"timestamp": time.Now().UTC().Format(time.RFC3339Nano),
		})
	})
}
