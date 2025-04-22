package service

import (
	"net"

	persistencev1 "github.com/purplepudding/bricks/api/pkg/pb/bricks/v1/persistence"
	"github.com/purplepudding/bricks/lib/microservice"
	"github.com/purplepudding/bricks/persistence/config"
	"github.com/purplepudding/bricks/persistence/internal/core/storage"
	"github.com/purplepudding/bricks/persistence/internal/grpcsvc"
	"github.com/purplepudding/bricks/persistence/internal/persistence"
	"google.golang.org/grpc"
)

var _ microservice.Service[*config.Config] = (*Service)(nil)

type Service struct {
	server *grpc.Server
	cfg    *config.Config
}

func (service *Service) Wire(cfg *config.Config) error {
	service.cfg = cfg

	st, err := persistence.NewNatsKVPersistence(cfg.NATS)
	if err != nil {
		return err
	}

	sl := storage.NewLogic(st)
	ss := grpcsvc.NewStorageService(sl)

	service.server = microservice.GRPCServer(func(g *grpc.Server) {
		persistencev1.RegisterStorageServiceServer(g, ss)
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
