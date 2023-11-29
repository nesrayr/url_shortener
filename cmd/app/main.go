package main

import (
	"context"
	"fmt"
	"github.com/heetch/confita"
	"github.com/heetch/confita/backend/env"
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
	"syscall"
	"url_shortener/config"
	"url_shortener/internal/migrations"
	"url_shortener/internal/storage/postgres"
	"url_shortener/pkg/logging"
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

	db, err := postgres.ConnectDB(cfg.Database)
	if err != nil {
		logger.Fatal(err)
	}
	err = db.Ping(ctx)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Info("successfully connected to db")

	migrations.MigrateDB("up", logger, cfg.Database, "./internal/migrations")

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
}
