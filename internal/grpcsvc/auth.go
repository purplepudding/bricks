package grpcsvc

import (
	"context"

	"buf.build/gen/go/purplepudding/foundation/grpc/go/foundation/v1/foundationv1grpc"
	"buf.build/gen/go/purplepudding/foundation/protocolbuffers/go/foundation/v1"
)

var _ foundationv1grpc.AuthServiceServer = (*AuthService)(nil)

type AuthService struct{}

func (a *AuthService) Login(ctx context.Context, req *foundationv1.LoginRequest) (*foundationv1.LoginResponse, error) {
	//TODO implement me
	panic("implement me")
}
