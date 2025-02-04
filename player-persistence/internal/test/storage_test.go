package test

import (
	"context"
	"testing"
	"time"

	"github.com/purplepudding/foundation/api/pkg/pb/foundation/v1/playerpersistence"
	"github.com/purplepudding/foundation/lib/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/structpb"
)

const (
	playerUUID = "d65fa999-3aea-4006-8aa2-3e7f0fb85334"
)

func TestIntegration_Set(t *testing.T) {
	val, err := structpb.NewStruct(map[string]any{"foo": "bar"})
	require.NoError(t, err)

	tests := []struct {
		name string
		req  *playerpersistence.SetRequest
		resp *playerpersistence.SetResponse
		code codes.Code
	}{
		{
			name: "can set some data",
			req: &playerpersistence.SetRequest{
				PlayerId: playerUUID,
				Datatype: "testData",
				Value:    val,
			},
			resp: &playerpersistence.SetResponse{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cc, err := grpc.NewClient("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
			require.NoError(t, err)
			defer func() {
				_ = cc.Close()
			}()

			cli := playerpersistence.NewStorageServiceClient(cc)

			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			resp, err := cli.Set(ctx, tt.req)
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

func TestIntegration_Get(t *testing.T) {
	val, err := structpb.NewStruct(map[string]any{"foo": "bar"})
	require.NoError(t, err)

	tests := []struct {
		name string
		req  *playerpersistence.GetRequest
		resp *playerpersistence.GetResponse
		code codes.Code
	}{
		{
			name: "can get some data",
			req: &playerpersistence.GetRequest{
				PlayerId: playerUUID,
				Datatype: "testData",
			},
			resp: &playerpersistence.GetResponse{
				Value: val,
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

			cli := playerpersistence.NewStorageServiceClient(cc)

			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			resp, err := cli.Get(ctx, tt.req)
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
