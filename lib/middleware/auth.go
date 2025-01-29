package middleware

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	HeaderAuthorizationKey = "Authorization"
)

func AuthUnaryInterceptor(ctx context.Context, req any, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	if err := intercept(ctx); err != nil {
		return nil, err
	}

	return handler(ctx, req)
}

func AuthStreamInterceptor(srv interface{}, ss grpc.ServerStream, _ *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	if err := intercept(ss.Context()); err != nil {
		return err
	}

	return handler(srv, ss)
}

func intercept(ctx context.Context) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Errorf(codes.Unauthenticated, "unauthenticated")
	}

	values := md.Get(HeaderAuthorizationKey)
	if len(values) == 0 {
		return status.Errorf(codes.Unauthenticated, "unauthenticated")
	}

	if len(values) > 1 {
		return status.Errorf(codes.Unauthenticated, "unauthenticated")
	}

	//TODO validate provided token

	return nil
}
