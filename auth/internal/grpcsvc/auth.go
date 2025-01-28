package grpcsvc

import (
	"context"
	"time"

	authv1 "github.com/purplepudding/foundation/api/pkg/pb/foundation/v1/auth"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ authv1.AuthServiceServer = (*AuthService)(nil)

type AuthService struct {
	authv1.UnimplementedAuthServiceServer
}

func (a *AuthService) Login(ctx context.Context, req *authv1.LoginRequest) (*authv1.LoginResponse, error) {
	return &authv1.LoginResponse{
		AccessToken:  "eyy.nice.token",
		RefreshToken: "eyy.refreshing.token",
		Expiry:       timestamppb.New(time.Now().Add(time.Hour)),
	}, nil
}
