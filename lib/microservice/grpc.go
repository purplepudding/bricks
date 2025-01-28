package microservice

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GRPCRegistrationFunc func(server *grpc.Server)

func GRPCServer(reg GRPCRegistrationFunc, opts ...grpc.ServerOption) *grpc.Server {
	g := grpc.NewServer(opts...)

	reflection.Register(g)
	//TODO add health check apis
	reg(g)

	return g
}
