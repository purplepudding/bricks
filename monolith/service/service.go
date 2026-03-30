package service

import (
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/nats-io/nats-server/v2/server"
	authSvc "github.com/purplepudding/bricks/auth/service"
	gatewayCfg "github.com/purplepudding/bricks/gateway/config"
	gatewaySvc "github.com/purplepudding/bricks/gateway/service"
	itemSvc "github.com/purplepudding/bricks/item/service"
	"github.com/purplepudding/bricks/lib/microservice"
	matchmakingSvc "github.com/purplepudding/bricks/matchmaking/service"
	"github.com/purplepudding/bricks/monolith/config"
	persistenceSvc "github.com/purplepudding/bricks/persistence/service"
	settingsSvc "github.com/purplepudding/bricks/settings/service"
	"golang.org/x/sync/errgroup"
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

	item := new(itemSvc.Service)
	if err := item.Wire(&cfg.Item); err != nil {
		return fmt.Errorf("failed to wire item service: %w", err)
	}
	service.servers["item"] = item

	matchmaking := new(matchmakingSvc.Service)
	if err := matchmaking.Wire(&cfg.Matchmaking); err != nil {
		return fmt.Errorf("failed to wire matchmaking service: %w", err)
	}
	service.servers["matchmaking"] = matchmaking

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

	gateway := new(gatewaySvc.Service)
	if err := gateway.Wire(&gatewayCfg.Config{
		Auth:         cfg.Auth,
		Item:         cfg.Item,
		Matchmaking:  cfg.Matchmaking,
		Persistence:  cfg.Persistence,
		Settings:     cfg.Settings,
		Microservice: cfg.Gateway,
	}); err != nil {
		return fmt.Errorf("failed to wire gateway service: %w", err)
	}
	service.servers["gateway"] = gateway

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
	var eg errgroup.Group

	for name, svc := range service.servers {
		//TODO use the errorgroup context to allow this to gracefully shut down
		eg.Go(func() error {
			if err := svc.Run(); err != nil {
				return fmt.Errorf("failed to boot %s service: %w", name, err)
			}
			return nil
		})
	}

	slog.Info("startup complete")

	// Wait for any errors to occur
	if err := eg.Wait(); err != nil {
		return err
	}

	return nil
}
