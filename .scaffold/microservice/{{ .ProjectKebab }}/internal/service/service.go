package service

import (
	"net"

	{{.ProjectKebab}}v1 "github.com/purplepudding/foundation/api/pkg/pb/foundation/v1/{{.ProjectKebab}}"
	"github.com/purplepudding/foundation/{{.ProjectKebab}}/internal/config"
	"github.com/purplepudding/foundation/{{.ProjectKebab}}/internal/grpcsvc"
	"github.com/purplepudding/foundation/lib/microservice"
	"google.golang.org/grpc"
)

var _ microservice.Service[*config.Config] = (*Service)(nil)

type Service struct {
	server *grpc.Server
}

func (service *Service) Wire(cfg *config.Config) error {
	service.server = microservice.GRPCServer(func(g *grpc.Server) {
		{{.ProjectKebab}}v1.RegisterAAAServiceServer(g, &grpcsvc.AAAService{})
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
