package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hifat/con-q-api/internal/app/config"
	"github.com/hifat/con-q-api/internal/app/di"
	"github.com/hifat/con-q-api/internal/app/routes/routeV1"
	"github.com/hifat/con-q-api/internal/pkg/validity"
)

func configCors() cors.Config {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = []string{
		"Origin",
		"Content-Length",
		"Content-Type",
		"Authorization",
	}

	return corsConfig
}

func main() {
	cfg := config.LoadAppConfig()
	wireHandler, cleanUp := di.InitializeAPI(cfg)
	defer cleanUp()

	validity.Register()

	router := gin.Default()
	router.Use(gin.Recovery())

	api := router.Group("")

	v1 := routeV1.New(api, wireHandler.Handler)
	v1.Register()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	srv := &http.Server{
		Addr:           cfg.Env.AppHost + ":" + cfg.Env.AppPort,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	<-ctx.Done()
	stop()
	log.Println("shutting down gracefully, press Ctrl+C again to force")

	timeOutctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(timeOutctx); err != nil {
		log.Println(err)
	}
}
