package main

import (
	"context"
	_ "github.com/lib/pq"
	"github.com/lifedaemon-kill/ozon-url-shortener-api/internal/pkg/config"
	"github.com/lifedaemon-kill/ozon-url-shortener-api/internal/pkg/lib"
	"github.com/lifedaemon-kill/ozon-url-shortener-api/internal/pkg/logger"
	"github.com/lifedaemon-kill/ozon-url-shortener-api/internal/server/grpc_server"
	"github.com/lifedaemon-kill/ozon-url-shortener-api/internal/server/http_server"
	"github.com/lifedaemon-kill/ozon-url-shortener-api/internal/service"
	"github.com/lifedaemon-kill/ozon-url-shortener-api/internal/storage"
	"github.com/lifedaemon-kill/ozon-url-shortener-api/internal/storage/inmemory"
	"github.com/lifedaemon-kill/ozon-url-shortener-api/internal/storage/postgres"
	"google.golang.org/grpc"
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

	urlService := service.New(conf.URLGenerator, repo, conf.ShortLinksAddress, log)
	grpcServer := grpc.NewServer()
	grpc_server.Registration(grpcServer, grpc_server.NewURLService(urlService, log))

	h := http_server.NewHandler(log, urlService)
	r := http_server.NewGinRouter(conf.ENV, h, log)

	httpServer := &http.Server{
		Addr:    conf.Http.Address,
		Handler: r,
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	//grpc
	go func() {
		<-ctx.Done()
		log.Info("shutting down API gRPC server")
		grpcServer.GracefulStop()
	}()

	//http
	go func() {
		if err = httpServer.ListenAndServe(); err != nil {
			log.Error("http server error", "error", err)
		}
	}()

	//graceful shd
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info("shutting down API server...")

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err = httpServer.Shutdown(ctx); err != nil {
		log.Error("shutdown server error", "error", err)
	}
	log.Info("shutdown complete")
}
