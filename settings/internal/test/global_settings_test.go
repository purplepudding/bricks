package test

import (
	"context"
	"testing"
	"time"

	"github.com/purplepudding/bricks/api/pkg/pb/bricks/v1/settings"
	"github.com/purplepudding/bricks/lib/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/structpb"
)

func TestIntegration_SetGlobalSettings(t *testing.T) {
	tests := []struct {
		name string
		req  *settings.SetGlobalSettingsRequest
		code codes.Code
	}{
		{
			name: "can set settings",
			req: &settings.SetGlobalSettingsRequest{
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

			cli := settings.NewGlobalSettingsServiceClient(cc)

			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			resp, err := cli.SetGlobalSettings(ctx, tt.req)
			if tt.code != codes.OK {
				assert.Nil(t, resp)
				assert.Error(t, err)

				s, _ := status.FromError(err)
				assert.Equal(t, tt.code, s.Code())
			} else {
				assert.NoError(t, err)
				test.ProtoEq(t, &settings.SetGlobalSettingsResponse{}, resp)
			}
		})
	}
}
