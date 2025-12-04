package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang.org/x/time/rate"

	"github.com/berjis/markets/backend/internal/cache"
	"github.com/berjis/markets/backend/internal/config"
	"github.com/berjis/markets/backend/internal/database"
	"github.com/berjis/markets/backend/internal/httpapi"
	"github.com/berjis/markets/backend/internal/middleware"
	"github.com/berjis/markets/backend/internal/migrations"
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

	rootCtx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	if err := migrations.Run(rootCtx, cfg.DatabaseURL); err != nil {
		log.Fatal("migration failed", zap.Error(err))
	}

	dbPool, err := database.NewPool(rootCtx, cfg.DatabaseURL, cfg.DBMaxConnections)
	if err != nil {
		log.Fatal("database connection failed", zap.Error(err))
	}
	defer dbPool.Close()

	cacheClient, err := cache.NewClient(cfg.RedisAddr, cfg.RedisPassword)
	if err != nil {
		log.Fatal("redis connection failed", zap.Error(err))
	}
	defer cacheClient.Close()

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middleware.RequestID())
	router.Use(middleware.RequestLogger(log))
	router.Use(middleware.RateLimiter(middleware.RateLimiterConfig{
		RequestsPerSecond: rate.Limit(100),
		Burst:             200,
	}))
	router.Use(cors.New(cors.Config{
		AllowOrigins:     cfg.AllowedOrigins(),
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Content-Type", "X-Request-Id"},
		ExposeHeaders:    []string{"X-Request-Id"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	httpapi.Register(router, httpapi.Dependencies{
		DB:     dbPool,
		Cache:  cacheClient,
		Logger: log,
		Config: cfg,
	})

	srv := server.NewHTTPServer(cfg, router)

	if err := srv.Run(rootCtx); err != nil {
		log.Error("http server shutdown", zap.Error(err))
		os.Exit(1)
	}
}
