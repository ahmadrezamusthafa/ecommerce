package rest

import (
	"context"
	"errors"
	"fmt"
	"github.com/ahmadrezamusthafa/ecommerce/config"
	"github.com/ahmadrezamusthafa/ecommerce/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Router struct {
	*gin.Engine
	cfg *config.Configuration
}

func InitRouter(cfg *config.Configuration) *Router {
	router := gin.Default()

	return &Router{
		Engine: router,
		cfg:    cfg,
	}
}

func (router *Router) Run() {
	port := router.cfg.App.Port
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router.Handler(),
	}

	logger.Infof("Starting http server on port %d", port)

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Fatalf("Failed to serve server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("Shutdown server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatalf("Server shutdown: %v", err)
	}
	select {
	case <-ctx.Done():
	}
	logger.Info("Server exiting")
}
