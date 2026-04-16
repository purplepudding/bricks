package service

import (
	"log/slog"
	"net"
	"net/netip"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/realip"
	matchmakingv1 "github.com/purplepudding/bricks/api/pkg/pb/bricks/v1/matchmaking"
	"github.com/purplepudding/bricks/lib/microservice"
	"github.com/purplepudding/bricks/matchmaking/config"
	"github.com/purplepudding/bricks/matchmaking/internal/core"
	"github.com/purplepudding/bricks/matchmaking/internal/grpcsvc"
	"github.com/purplepudding/bricks/matchmaking/internal/memorymatch"
	"google.golang.org/grpc"
)

var _ microservice.Service[*config.Config] = (*Service)(nil)

type Service struct {
	server *grpc.Server
	cfg    *config.Config
}

func (service *Service) Wire(cfg *config.Config) error {
	service.cfg = cfg

	memoryMatchmakingClient := memorymatch.NewMemoryMatchClient()
	matchmaker := core.NewMatchmaker(memoryMatchmakingClient)

	trustedPeers := []netip.Prefix{netip.MustParsePrefix("127.0.0.1/32"), netip.MustParsePrefix("10.0.0.0/8")}
	headers := []string{realip.XForwardedFor}

	service.server = microservice.GRPCServer(cfg, func(g *grpc.Server) {
		matchmakingv1.RegisterMatchmakingServiceServer(g, grpcsvc.NewMatchmakingService(matchmaker))
	},
		grpc.ChainUnaryInterceptor(realip.UnaryServerInterceptorOpts(realip.WithTrustedPeers(trustedPeers), realip.WithHeaders(headers), realip.WithTrustedProxiesCount(1))),
		grpc.ChainStreamInterceptor(realip.StreamServerInterceptorOpts(realip.WithTrustedPeers(trustedPeers), realip.WithHeaders(headers), realip.WithTrustedProxiesCount(1))),
	)

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
