package service

import (
	"net"

	"buf.build/gen/go/purplepudding/foundation/grpc/go/foundation/v1/foundationv1grpc"
	"github.com/purplepudding/foundation/internal/grpcsvc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Service struct {
	server *grpc.Server
}

func NewService() *Service {
	//TODO load config from file
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	reflection.Register(grpcServer)
	foundationv1grpc.RegisterAuthServiceServer(grpcServer, &grpcsvc.AuthService{})

	return &Service{
		server: grpcServer,
	}
}

func (service *Service) Run() {
	//TODO get address from config

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	if err := service.server.Serve(lis); err != nil {
		panic(err)
	}
	//TODO handle exit signals
}
