package main

import (
	"Backend/admin_panel/internal/config"
	"Backend/admin_panel/internal/domain/user"
	"Backend/admin_panel/internal/http/handlers"
	"Backend/admin_panel/internal/http/router"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	pmongo "Backend/admin_panel/internal/platform/mongo"

	"go.uber.org/zap"
)

func main() {
	cfg := config.LoadConfig()
	logr, _ := zap.NewProduction()
	defer logr.Sync()

	//Mongo
	mc, err := pmongo.New(cfg.MongoURI, cfg.MongoDB, logr)
	if err != nil {
		log.Fatalf("mongo connect: %v", err)
	}

	//DI
	repo := user.NewRepository(mc.DB)
	if err := repo.EnsureIndexes(context.Background()); err != nil {
		log.Fatalf("ensure indexes: %v", err)
	}
	svc := user.NewService(repo)
	userHandler := handlers.NewUserHandler(svc)
	router := router.New(userHandler)

	srv := &http.Server{Addr: ":" + cfg.AppPort, Handler: router, ReadHeaderTimeout: 10 * time.Second}

	go func() {
		logr.Info("server starting", zap.String("port", cfg.AppPort))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_ = srv.Shutdown(ctx)
	logr.Info("server stopped")

}
