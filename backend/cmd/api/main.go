package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang.org/x/time/rate"

	"github.com/berjis/markets/backend/internal/config"
	"github.com/berjis/markets/backend/internal/httpapi"
	"github.com/berjis/markets/backend/internal/middleware"
	"github.com/berjis/markets/backend/internal/server"
	"github.com/berjis/markets/backend/pkg/logger"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	log, err := logger.New(cfg.LogLevel)
	if err != nil {
		panic(err)
	}
	defer logger.Sync(log)

	if cfg.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middleware.RequestID())
	router.Use(middleware.RequestLogger(log))
	router.Use(middleware.RateLimiter(middleware.RateLimiterConfig{
		RequestsPerSecond: rate.Limit(100),
		Burst:             200,
	}))

	httpapi.Register(router)

	srv := server.NewHTTPServer(cfg, router)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	if err := srv.Run(ctx); err != nil {
		log.Error("http server shutdown", zap.Error(err))
		os.Exit(1)
	}
}
