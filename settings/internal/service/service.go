package service

import (
	"net"

	settingsv1 "github.com/purplepudding/foundation/api/pkg/pb/foundation/v1/settings"
	"github.com/purplepudding/foundation/lib/microservice"
	"github.com/purplepudding/foundation/settings/internal/config"
	"github.com/purplepudding/foundation/settings/internal/core/settings"
	v1 "github.com/purplepudding/foundation/settings/internal/grpcsvc/v1"
	"github.com/purplepudding/foundation/settings/internal/persistence"
	"github.com/valkey-io/valkey-go"
	"google.golang.org/grpc"
)

var _ microservice.Service[*config.Config] = (*Service)(nil)

type Service struct {
	server *grpc.Server
}

func (service *Service) Wire(cfg *config.Config) error {
	//TODO add some backoff to this so we don't immediately fail if Valkey can't be accessed - only mark unready
	valkeyCli, err := valkey.NewClient(valkey.ClientOption{InitAddress: []string{cfg.Valkey.Addr}})
	if err != nil {
		return err //TODO sentinel or wrap
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
	//TODO get address from config
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		return err
	}

	if err := service.server.Serve(lis); err != nil {
		return err
	}

	return nil
}
