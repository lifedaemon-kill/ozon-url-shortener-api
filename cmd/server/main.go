package main

import (
	"context"
	"github.com/lifedaemon-kill/ozon-url-shortener-api/internal/pkg/config"
	"github.com/lifedaemon-kill/ozon-url-shortener-api/internal/pkg/lib"
	"github.com/lifedaemon-kill/ozon-url-shortener-api/internal/pkg/logger"
	"github.com/lifedaemon-kill/ozon-url-shortener-api/internal/server/http_server"
	"github.com/lifedaemon-kill/ozon-url-shortener-api/internal/service"
	"github.com/lifedaemon-kill/ozon-url-shortener-api/internal/storage"
	"github.com/lifedaemon-kill/ozon-url-shortener-api/internal/storage/inmemory"
	"github.com/lifedaemon-kill/ozon-url-shortener-api/internal/storage/postgres"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const configPath = "config/config.yaml"

func main() {
	conf := config.Load(configPath)
	log := logger.SetUpLogger(conf.ENV)

	storageType, err := lib.ParseStorageType(os.Args[1:])
	if err != nil {
		log.Error("wrong params", "err", err)
		return
	}

	var repo storage.Storage

	switch storageType {
	case lib.Postgres:
		repo, err = postgres.New(conf.DB)
		if err != nil {
			log.Error("ошибка при подключении к postgres", "err", err, "conf", conf.DB)
			return
		}
		log.Info("repository created", "type", lib.Postgres)
	case lib.InMemory:
		repo = inmemory.New()
	default:
		log.Error("unknown storage type", "storageType", storageType, "expected", lib.Postgres+" or "+lib.InMemory)
		log.Info("repository created", "type", lib.InMemory)
		return
	}

	urlService := service.New(conf.URLGenerator, repo, log)
	h := http_server.NewHandler(log, urlService)
	r := http_server.NewGinRouter(conf.ENV, h, log)

	httpServer := &http.Server{
		Addr:    conf.Http.Address,
		Handler: r,
	}

	go func() {
		if err = httpServer.ListenAndServe(); err != nil {
			log.Error("http server error", "error", err)
			os.Exit(1)
		}
	}()

	//graceful shd
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = httpServer.Shutdown(ctx); err != nil {
		log.Error("shutdown server error", "error", err)
	}
}
