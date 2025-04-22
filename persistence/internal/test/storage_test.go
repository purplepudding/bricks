package test

import (
	"context"
	"testing"
	"time"

	persistencev1 "github.com/purplepudding/bricks/api/pkg/pb/bricks/v1/persistence"
	"github.com/purplepudding/bricks/lib/test"
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
	titleUUID  = "a123fa999-3aea-4006-8aa2-3e7f0fb85334"
)

func TestIntegration_Set(t *testing.T) {
	val, err := structpb.NewStruct(map[string]any{"foo": "bar"})
	require.NoError(t, err)

	tests := []struct {
		name string
		req  *persistencev1.SetRequest
		resp *persistencev1.SetResponse
		code codes.Code
	}{
		{
			name: "can set some player data",
			req: &persistencev1.SetRequest{
				Key: &persistencev1.Key{
					TypedKey: &persistencev1.Key_PlayerKey{
						PlayerKey: &persistencev1.PlayerKey{
							TitleId:  titleUUID,
							PlayerId: playerUUID,
							Datatype: "testData",
						},
					},
				},
				Value: val,
			},
			resp: &persistencev1.SetResponse{},
		},
		{
			name: "can set some title data",
			req: &persistencev1.SetRequest{
				Key: &persistencev1.Key{
					TypedKey: &persistencev1.Key_TitleKey{
						TitleKey: &persistencev1.TitleKey{
							TitleId:  titleUUID,
							Datatype: "testData",
						},
					},
				},
				Value: val,
			},
			resp: &persistencev1.SetResponse{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cc, err := grpc.NewClient("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
			require.NoError(t, err)
			defer func() {
				_ = cc.Close()
			}()

			cli := persistencev1.NewStorageServiceClient(cc)

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
		req  *persistencev1.GetRequest
		resp *persistencev1.GetResponse
		code codes.Code
	}{
		{
			name: "can get some player data",
			req: &persistencev1.GetRequest{
				Key: &persistencev1.Key{
					TypedKey: &persistencev1.Key_PlayerKey{
						PlayerKey: &persistencev1.PlayerKey{
							TitleId:  titleUUID,
							PlayerId: playerUUID,
							Datatype: "testData",
						},
					},
				},
			},
			resp: &persistencev1.GetResponse{
				Value: val,
			},
		},
		{
			name: "can get some title data",
			req: &persistencev1.GetRequest{
				Key: &persistencev1.Key{
					TypedKey: &persistencev1.Key_TitleKey{
						TitleKey: &persistencev1.TitleKey{
							TitleId:  titleUUID,
							Datatype: "testData",
						},
					},
				},
			},
			resp: &persistencev1.GetResponse{
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

			cli := persistencev1.NewStorageServiceClient(cc)

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
