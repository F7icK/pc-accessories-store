package main

import (
	"context"
	"errors"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/F7icK/pc-accessories-store/internal/clients/repository"
	"github.com/F7icK/pc-accessories-store/internal/server"
	"github.com/F7icK/pc-accessories-store/internal/server/handlers"
	"github.com/F7icK/pc-accessories-store/internal/service"
	"github.com/F7icK/pc-accessories-store/internal/types/config"
	"github.com/F7icK/pc-accessories-store/pkg/database/postgres"
	"github.com/labstack/gommon/log"
	"gopkg.in/yaml.v3"
)

func main() {
	configPath := new(string)
	flag.StringVar(configPath, "config-path", "./config/config-default.yaml", "specify path to yaml")
	flag.Parse()

	configFile, err := os.Open(*configPath)
	if err != nil {
		log.Errorf("err with os.Open config: %s", err)
	}

	cfg := config.Config{}
	if err = yaml.NewDecoder(configFile).Decode(&cfg); err != nil {
		log.Errorf("err with Decode config: %s", err)
	}

	postgresClient, err := postgres.NewPostgres(cfg.PostgresDsn)
	if err != nil {
		log.Fatalf("err with NewPostgres: %s", err)
	}

	db, err := postgresClient.Database()
	if err != nil {
		log.Fatalf("err with Gorm: %s", err)
	}

	repos := repository.NewRepository(db)

	services := service.NewService(repos)

	endpoints := handlers.NewHandlers(services)

	srv := server.NewServer(&cfg, endpoints)

	stopFunc := func() {
		if err = srv.Shutdown(context.Background()); err != nil {
			log.Fatalf("failed to stop server: %s", err)
		}

		if err = postgresClient.Close(); err != nil {
			log.Fatalf("failed to stop db: %s", err)
		}
	}

	signalCh := make(chan os.Signal)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
	defer close(signalCh)

	go func(signalCh <-chan os.Signal, stopFunc func()) {
		select {
		case sig := <-signalCh:
			log.Infof("stopped with signal: %s", sig)
			stopFunc()
			os.Exit(0)
		}
	}(signalCh, stopFunc)

	if err = srv.Run(); !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("err with NewRouter: %s", err)
	}

}
