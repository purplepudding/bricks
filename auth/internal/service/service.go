package service

import (
	"net"

	authv1 "github.com/purplepudding/foundation/api/pkg/pb/foundation/v1/auth"
	"github.com/purplepudding/foundation/auth/internal/config"
	"github.com/purplepudding/foundation/auth/internal/grpcsvc"
	"github.com/purplepudding/foundation/lib/microservice"
	"google.golang.org/grpc"
)

var _ microservice.Service[*config.Config] = (*Service)(nil)

type Service struct {
	server *grpc.Server
}

func (service *Service) Wire(_ *config.Config) error {
	service.server = microservice.GRPCServer(func(g *grpc.Server) {
		authv1.RegisterAuthServiceServer(g, &grpcsvc.AuthService{})
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
