package service

import (
	"net"

	authv1 "github.com/purplepudding/foundation/api/pkg/pb/foundation/v1/auth"
	"github.com/purplepudding/foundation/auth/config"
	"github.com/purplepudding/foundation/auth/internal/grpcsvc"
	"github.com/purplepudding/foundation/lib/microservice"
	"google.golang.org/grpc"
)

var _ microservice.Service[*config.Config] = (*Service)(nil)

type Service struct {
	server *grpc.Server
	cfg    *config.Config
}

func (service *Service) Wire(cfg *config.Config) error {
	service.server = microservice.GRPCServer(func(g *grpc.Server) {
		authv1.RegisterAuthServiceServer(g, &grpcsvc.AuthService{})
	})
	service.cfg = cfg

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
