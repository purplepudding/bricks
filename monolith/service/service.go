package service

import (
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/nats-io/nats-server/v2/server"
	authSvc "github.com/purplepudding/foundation/auth/service"
	"github.com/purplepudding/foundation/lib/microservice"
	"github.com/purplepudding/foundation/monolith/config"
	persistenceSvc "github.com/purplepudding/foundation/persistence/service"
	settingsSvc "github.com/purplepudding/foundation/settings/service"
)

var _ microservice.Service[*config.Config] = (*Service)(nil)

type Service struct {
	nats    *server.Server
	servers map[string]microservice.Runnable
}

func (service *Service) Wire(cfg *config.Config) error {
	if err := service.startNATS(); err != nil {
		return err
	}

	service.servers = make(map[string]microservice.Runnable)

	auth := new(authSvc.Service)
	if err := auth.Wire(&cfg.Auth); err != nil {
		return fmt.Errorf("failed to wire auth service: %w", err)
	}
	service.servers["auth"] = auth

	persistence := new(persistenceSvc.Service)
	if err := persistence.Wire(&cfg.Persistence); err != nil {
		return fmt.Errorf("failed to wire persistence service: %w", err)
	}
	service.servers["persistence"] = persistence

	settings := new(settingsSvc.Service)
	if err := settings.Wire(&cfg.Settings); err != nil {
		return fmt.Errorf("failed to wire settings service: %w", err)
	}
	service.servers["settings"] = settings

	return nil
}

func (service *Service) startNATS() error {
	opts := &server.Options{
		Host:      "localhost",
		Port:      4222,
		JetStream: true,
	}
	nats, err := server.NewServer(opts)
	if err != nil {
		return fmt.Errorf("could not create nats server: %w", err)
	}
	service.nats = nats

	slog.Info("starting embedded NATS")

	go service.nats.Start()

	if !service.nats.ReadyForConnections(5 * time.Second) {
		return errors.New("nats server startup timed out")
	}

	slog.Info("embedded NATS server started")
	return nil
}

func (service *Service) Run() error {
	for name, svc := range service.servers {
		slog.Info("starting service", "svc", name)
		if err := svc.Run(); err != nil {
			return fmt.Errorf("failed to boot %s service: %w", name, err)
		}
	}

	slog.Info("startup complete")

	return nil
}
