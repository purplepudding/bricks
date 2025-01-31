package test

import (
	"context"
	"testing"
	"time"

	"github.com/purplepudding/foundation/api/pkg/pb/foundation/v1/settings"
	"github.com/purplepudding/foundation/lib/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/structpb"
)

func TestIntegration_GetServiceSettings(t *testing.T) {
	service := "service"

	tests := []struct {
		name string
		req  *settings.GetServiceSettingsRequest
		resp *settings.GetServiceSettingsResponse
		code codes.Code
	}{
		{
			name: "can get service settings",
			req: &settings.GetServiceSettingsRequest{
				Service: service,
			},
			resp: &settings.GetServiceSettingsResponse{
				Settings: map[string]*structpb.Value{
					"foo": structpb.NewStringValue("bar"),
				},
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

			cli := settings.NewServiceSettingsServiceClient(cc)

			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			resp, err := cli.GetServiceSettings(ctx, tt.req)
			if tt.code != codes.OK {
				assert.Nil(t, resp)
				assert.Error(t, err)

				s, _ := status.FromError(err)
				assert.Equal(t, tt.code, s.Code())
			} else {
				assert.NoError(t, err)
				test.ProtoEq(t, tt.resp, resp)
			}
		})
	}
}

func TestIntegration_SetServiceSettings(t *testing.T) {
	service := "service"

	tests := []struct {
		name string
		req  *settings.SetServiceSettingsRequest
		code codes.Code
	}{
		{
			name: "can set service settings",
			req: &settings.SetServiceSettingsRequest{
				Service: service,
				Settings: map[string]*structpb.Value{
					"baz": structpb.NewStringValue("bat"),
				},
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

			cli := settings.NewServiceSettingsServiceClient(cc)

			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			resp, err := cli.SetServiceSettings(ctx, tt.req)
			if tt.code != codes.OK {
				assert.Nil(t, resp)
				assert.Error(t, err)

				s, _ := status.FromError(err)
				assert.Equal(t, tt.code, s.Code())
			} else {
				assert.NoError(t, err)
				test.ProtoEq(t, &settings.SetServiceSettingsResponse{}, resp)
			}
		})
	}
}
