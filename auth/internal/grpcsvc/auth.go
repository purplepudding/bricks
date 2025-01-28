package grpcsvc

import (
	"context"

	authv1 "github.com/purplepudding/foundation/api/pkg/pb/foundation/v1/auth"
)

var _ authv1.AuthServiceServer = (*AuthService)(nil)

type AuthService struct {
	authv1.UnimplementedAuthServiceServer
}

func (a *AuthService) Login(ctx context.Context, req *authv1.LoginRequest) (*authv1.LoginResponse, error) {
	//TODO implement me
	panic("implement me")
}
