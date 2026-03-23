package service

import (
	"log/slog"
	"net"

	matchmakingv1 "github.com/purplepudding/bricks/api/pkg/pb/bricks/v1/matchmaking"
	"github.com/purplepudding/bricks/lib/microservice"
	"github.com/purplepudding/bricks/matchmaking/config"
	"github.com/purplepudding/bricks/matchmaking/internal/core"
	"github.com/purplepudding/bricks/matchmaking/internal/grpcsvc"
	"google.golang.org/grpc"
)

var _ microservice.Service[*config.Config] = (*Service)(nil)

type Service struct {
	server *grpc.Server
	cfg    *config.Config
}

func (service *Service) Wire(cfg *config.Config) error {
	service.cfg = cfg

	matchmaker := core.NewMatchmaker()

	service.server = microservice.GRPCServer(func(g *grpc.Server) {
		matchmakingv1.RegisterMatchmakingServiceServer(g, grpcsvc.NewMatchmakingService(matchmaker))
	})

	return nil
}

func (service *Service) Run() error {
	slog.Info("starting service", "svc", "matchmaking", "addr", service.cfg.ServingAddr)

	lis, err := net.Listen("tcp", service.cfg.ServingAddr)
	if err != nil {
		return err
	}

	if err := service.server.Serve(lis); err != nil {
		return err
	}

	return nil
}
