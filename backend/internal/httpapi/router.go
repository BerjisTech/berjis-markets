package httpapi

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"

	"github.com/berjis/markets/backend/internal/config"
	"github.com/berjis/markets/backend/internal/httpapi/routes"
)

// Dependencies bundles shared services for HTTP handlers.
type Dependencies struct {
	DB     DatabasePinger
	Cache  CachePinger
	Logger *zap.Logger
	Config config.Config
}

type DatabasePinger interface {
	Ping(context.Context) error
}

type CachePinger interface {
	Ping(context.Context) *redis.StatusCmd
}

// Register attaches all HTTP routes to the router.
func Register(router *gin.Engine, deps Dependencies) {
	routes.RegisterHealth(router, deps.DB, deps.Cache)
}
