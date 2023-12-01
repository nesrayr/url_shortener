package main

import (
	"context"
	"fmt"
	"github.com/heetch/confita"
	"github.com/heetch/confita/backend/env"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"url_shortener/config"
	"url_shortener/internal/migrations"
	grpc_ "url_shortener/internal/ports/grpc"
	router "url_shortener/internal/ports/http"
	"url_shortener/internal/repo"
	cacheRepo "url_shortener/internal/repo/cache"
	postgresRepo "url_shortener/internal/repo/postgres"
	"url_shortener/internal/service"
	"url_shortener/internal/storage/cache"
	"url_shortener/internal/storage/postgres"
	"url_shortener/pkg/logging"
)

const (
	postgresStorageMode = "postgres"
	cacheStorageMode    = "cache"
	httpTransportMode   = "http"
	grpcTransportMode   = "grpc"
)

func main() {
	logger := logging.GetLogger()

	ctx := context.Background()

	var cfg config.Config
	err := confita.NewLoader(
		env.NewBackend(),
	).Load(ctx, &cfg)
	if err != nil {
		logger.Fatal(err)
	}

	var repository repo.Repository
	if cfg.StorageMode == postgresStorageMode {
		db, err := postgres.ConnectDB(cfg.Database)
		if err != nil {
			logger.Fatal(err)
		}
		err = db.Ping(ctx)
		if err != nil {
			logger.Fatal(err)
		}
		defer db.Close()

		logger.Info("successfully connected to db")

		migrations.MigrateDB("up", logger, cfg.Database, "./internal/migrations")

		repository = postgresRepo.NewRepository(db, logger)
	} else if cfg.StorageMode == cacheStorageMode {
		mu := sync.RWMutex{}
		c := cache.NewCache(&mu)

		repository = cacheRepo.NewRepository(&c, logger)

		logger.Info("successfully created cache")
	}

	// configuring graceful shutdown
	sigQuit := make(chan os.Signal, 1)
	defer close(sigQuit)
	signal.Ignore(syscall.SIGHUP, syscall.SIGPIPE)
	signal.Notify(sigQuit, syscall.SIGINT, syscall.SIGTERM)

	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		select {
		case s := <-sigQuit:
			return fmt.Errorf("captured signal: %v", s)
		case <-ctx.Done():
			return nil
		}
	})

	serv := service.NewService(repository, logger)

	if cfg.TransportMode == httpTransportMode {
		r := router.SetupRouter(serv, logger)

		err = http.ListenAndServe(fmt.Sprintf(":%s", cfg.Port), r)
		if err != nil {
			logger.Fatal(err)
		}
	} else if cfg.TransportMode == grpcTransportMode {
		h := grpc_.NewHandler(serv, logger)
		srv := grpc_.NewServer(h, logger)

		err = srv.Run(cfg.Port)
		if err != nil {
			logger.Fatal(err)
		}
	}
}
