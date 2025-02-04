package service

import (
	"net"

	playerpersistencev1 "github.com/purplepudding/foundation/api/pkg/pb/foundation/v1/playerpersistence"
	"github.com/purplepudding/foundation/lib/microservice"
	"github.com/purplepudding/foundation/lib/valkeycli"
	"github.com/purplepudding/foundation/player-persistence/internal/config"
	"github.com/purplepudding/foundation/player-persistence/internal/core/storage"
	"github.com/purplepudding/foundation/player-persistence/internal/grpcsvc"
	"github.com/purplepudding/foundation/player-persistence/internal/persistence"
	"google.golang.org/grpc"
)

var _ microservice.Service[*config.Config] = (*Service)(nil)

type Service struct {
	server *grpc.Server
}

func (service *Service) Wire(cfg *config.Config) error {
	valkeyCli, err := valkeycli.New(cfg.Valkey)
	if err != nil {
		return err
	}

	st := persistence.NewValkeyStorage(valkeyCli)
	sl := storage.NewLogic(st)
	ss := grpcsvc.NewStorageService(sl)

	service.server = microservice.GRPCServer(func(g *grpc.Server) {
		playerpersistencev1.RegisterStorageServiceServer(g, ss)
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
