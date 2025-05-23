package service

import (
	"log/slog"
	"net"

	{{.ProjectKebab}}v1 "github.com/purplepudding/bricks/api/pkg/pb/bricks/v1/{{.ProjectKebab}}"
	"github.com/purplepudding/bricks/{{.ProjectKebab}}/config"
	"github.com/purplepudding/bricks/{{.ProjectKebab}}/internal/grpcsvc"
	"github.com/purplepudding/bricks/lib/microservice"
	"google.golang.org/grpc"
)

var _ microservice.Service[*config.Config] = (*Service)(nil)

type Service struct {
	server *grpc.Server
	cfg    *config.Config
}

func (service *Service) Wire(cfg *config.Config) error {
	service.cfg = cfg

	service.server = microservice.GRPCServer(func(g *grpc.Server) {
		{{.ProjectKebab}}v1.RegisterAAAServiceServer(g, &grpcsvc.AAAService{})
	})

	return nil
}

func (service *Service) Run() error {
	slog.Info("starting service", "svc", "{{.ProjectKebab}}", "addr", service.cfg.ServingAddr)

	lis, err := net.Listen("tcp", service.cfg.ServingAddr)
	if err != nil {
		return err
	}

	if err := service.server.Serve(lis); err != nil {
		return err
	}

	return nil
}
