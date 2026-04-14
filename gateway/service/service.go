package service

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	authv1 "github.com/purplepudding/bricks/api/pkg/pb/bricks/v1/auth"
	itemv1 "github.com/purplepudding/bricks/api/pkg/pb/bricks/v1/item"
	matchmakingv1 "github.com/purplepudding/bricks/api/pkg/pb/bricks/v1/matchmaking"
	persistencev1 "github.com/purplepudding/bricks/api/pkg/pb/bricks/v1/persistence"
	settingsv1 "github.com/purplepudding/bricks/api/pkg/pb/bricks/v1/settings"
	"github.com/purplepudding/bricks/gateway/config"
	"github.com/purplepudding/bricks/lib/clients/grpccli"
	"github.com/purplepudding/bricks/lib/microservice"
)

var _ microservice.Service[*config.Config] = (*Service)(nil)

type Service struct {
	mux *runtime.ServeMux
	cfg *config.Config
}

func (service *Service) Wire(cfg *config.Config) error {
	service.cfg = cfg
	service.mux = runtime.NewServeMux()

	dialOpts := grpccli.DialOpts()

	ctx := context.Background()

	// Auth
	if err := authv1.RegisterAuthServiceHandlerFromEndpoint(ctx, service.mux, cfg.Auth.ServingAddr, dialOpts); err != nil {
		return err
	}

	// Item
	if err := itemv1.RegisterCatalogServiceHandlerFromEndpoint(ctx, service.mux, cfg.Item.ServingAddr, dialOpts); err != nil {
		return err
	}

	// Matchmaking
	if err := matchmakingv1.RegisterMatchmakingServiceHandlerFromEndpoint(ctx, service.mux, cfg.Matchmaking.ServingAddr, dialOpts); err != nil {
		return err
	}

	// Persistence
	if err := persistencev1.RegisterStorageServiceHandlerFromEndpoint(ctx, service.mux, cfg.Persistence.ServingAddr, dialOpts); err != nil {
		return err
	}

	// Settings
	if err := settingsv1.RegisterGlobalSettingsServiceHandlerFromEndpoint(ctx, service.mux, cfg.Settings.ServingAddr, dialOpts); err != nil {
		return err
	}
	if err := settingsv1.RegisterItemParametersServiceHandlerFromEndpoint(ctx, service.mux, cfg.Settings.ServingAddr, dialOpts); err != nil {
		return err
	}
	if err := settingsv1.RegisterServiceSettingsServiceHandlerFromEndpoint(ctx, service.mux, cfg.Settings.ServingAddr, dialOpts); err != nil {
		return err
	}

	return nil
}

func (service *Service) Run() error {
	slog.Info("starting service", "svc", "gateway", "addr", service.cfg.ServingAddr)

	if err := http.ListenAndServe(service.cfg.ServingAddr, service.mux); err != nil {
		return err
	}

	return nil
}
