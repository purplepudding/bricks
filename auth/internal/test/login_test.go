package test

import (
	"context"
	"testing"
	"time"

	"github.com/purplepudding/bricks/api/pkg/pb/foundation/v1/auth"
	"github.com/purplepudding/bricks/lib/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/testing/protocmp"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestIntegration_Login(t *testing.T) {
	tests := []struct {
		name string
		req  *auth.LoginRequest
		resp *auth.LoginResponse
		code codes.Code
	}{
		{
			name: "can login with valid credentials",
			req: &auth.LoginRequest{
				Credentials: &auth.LoginRequest_UsernamePassword{
					UsernamePassword: &auth.UserPass{
						Username: "foo",
						Password: "bar",
					},
				},
			},
			resp: &auth.LoginResponse{
				AccessToken:  "eyy.nice.token",
				RefreshToken: "eyy.refreshing.token",
				Expiry:       timestamppb.New(time.Now().Add(time.Hour)), //this isn't checked
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cc, err := grpc.NewClient("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
			require.NoError(t, err)
			defer func() {
				_ = cc.Close()
			}()

			cli := auth.NewAuthServiceClient(cc)

			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			resp, err := cli.Login(ctx, tt.req)
			if tt.code != codes.OK {
				assert.Nil(t, resp)
				assert.Error(t, err)

				s, _ := status.FromError(err)
				assert.Equal(t, tt.code, s.Code())
			} else {
				assert.NoError(t, err)
				test.ProtoEq(t, tt.resp, resp, protocmp.IgnoreFields(&auth.LoginResponse{}, "expiry"))
			}
		})
	}
}
