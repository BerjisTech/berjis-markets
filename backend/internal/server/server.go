package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/berjis/markets/backend/internal/config"
)

// HTTPServer wraps the Gin engine with graceful shutdown primitives.
type HTTPServer struct {
	engine *gin.Engine
	srv    *http.Server
	cfg    config.Config
}

func NewHTTPServer(cfg config.Config, engine *gin.Engine) *HTTPServer {
	srv := &http.Server{
		Addr:              fmt.Sprintf(":%d", cfg.HTTPPort),
		Handler:           engine,
		ReadHeaderTimeout: cfg.ReadHeaderTimeout,
	}

	return &HTTPServer{engine: engine, srv: srv, cfg: cfg}
}

func (s *HTTPServer) Run(ctx context.Context) error {
	errCh := make(chan error, 1)

	go func() {
		if err := s.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errCh <- err
		}
	}()

	select {
	case <-ctx.Done():
		shutdownCtx, cancel := context.WithTimeout(context.Background(), s.cfg.ShutdownTimeout)
		defer cancel()
		return s.srv.Shutdown(shutdownCtx)
	case err := <-errCh:
		return err
	}
}

func (s *HTTPServer) Engine() *gin.Engine {
	return s.engine
}
