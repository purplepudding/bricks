package service

import (
	"log/slog"
	"net"

	itemv1 "github.com/purplepudding/bricks/api/pkg/pb/bricks/v1/item"
	"github.com/purplepudding/bricks/item/config"
	"github.com/purplepudding/bricks/item/internal/assetbundles"
	"github.com/purplepudding/bricks/item/internal/core/catalog"
	"github.com/purplepudding/bricks/item/internal/grpcsvc"
	"github.com/purplepudding/bricks/item/internal/persistence"
	"github.com/purplepudding/bricks/item/internal/settings"
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

	ip, err := persistence.New(cfg.Clients.NATS)
	if err != nil {
		//the nats client does retrying so at this point we should die and let K8s restart the pod
		panic(err)
	}

	abc := assetbundles.NewClient()
	isc := settings.NewClient(cfg.Clients.Settings)

	cl := catalog.New(ip, abc, isc)
	cs := grpcsvc.NewCatalogService(cl)

	service.server = microservice.GRPCServer(func(g *grpc.Server) {
		itemv1.RegisterCatalogServiceServer(g, cs)
	})

	return nil
}

func (service *Service) Run() error {
	slog.Info("starting service", "svc", "item", "addr", service.cfg.ServingAddr)

	lis, err := net.Listen("tcp", service.cfg.ServingAddr)
	if err != nil {
		return err
	}

	if err := service.server.Serve(lis); err != nil {
		return err
	}

	return nil
}
