package service

import (
	"net"

	settingsv1 "github.com/purplepudding/foundation/api/pkg/pb/foundation/v1/settings"
	"github.com/purplepudding/foundation/lib/microservice"
	"github.com/purplepudding/foundation/settings/internal/grpcsvc"
	"google.golang.org/grpc"
)

var _ microservice.Service = (*Service)(nil)

type Service struct {
	server *grpc.Server
}

func (service *Service) Wire() error {
	service.server = microservice.GRPCServer(func(g *grpc.Server) {
		settingsv1.RegisterGlobalSettingsServiceServer(g, &grpcsvc.GlobalSettingsService{})
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
