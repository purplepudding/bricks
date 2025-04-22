package service

import (
	"net"

	settingsv1 "github.com/purplepudding/bricks/api/pkg/pb/bricks/v1/settings"
	"github.com/purplepudding/bricks/lib/clients/valkeycli"
	"github.com/purplepudding/bricks/lib/microservice"
	"github.com/purplepudding/bricks/settings/config"
	"github.com/purplepudding/bricks/settings/internal/core/settings"
	v1 "github.com/purplepudding/bricks/settings/internal/grpcsvc/v1"
	"github.com/purplepudding/bricks/settings/internal/persistence"
	"google.golang.org/grpc"
)

var _ microservice.Service[*config.Config] = (*Service)(nil)

type Service struct {
	server *grpc.Server
	cfg    *config.Config
}

func (service *Service) Wire(cfg *config.Config) error {
	service.cfg = cfg

	valkeyCli, err := valkeycli.New(cfg.Valkey)
	if err != nil {
		return err
	}

	settingsStore := persistence.NewValkeySettingsStore(valkeyCli)

	gsLogic := settings.NewGlobalSettingsLogic(settingsStore)
	gsSvc := v1.NewGlobalSettingsService(gsLogic)

	ssLogic := settings.NewServiceSettingsLogic(gsLogic, settingsStore)
	ssSvc := v1.NewServiceSettingsService(ssLogic)

	service.server = microservice.GRPCServer(func(g *grpc.Server) {
		settingsv1.RegisterGlobalSettingsServiceServer(g, gsSvc)
		settingsv1.RegisterServiceSettingsServiceServer(g, ssSvc)
	})

	return nil
}

func (service *Service) Run() error {
	lis, err := net.Listen("tcp", service.cfg.ServingAddr)
	if err != nil {
		return err
	}

	if err := service.server.Serve(lis); err != nil {
		return err
	}

	return nil
}
