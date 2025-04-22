package test

import (
	"context"
	"testing"
	"time"

	"github.com/purplepudding/bricks/api/pkg/pb/foundation/v1/{{.ProjectKebab}}"
	"github.com/purplepudding/bricks/lib/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func TestIntegration_Do(t *testing.T) {
	tests := []struct {
		name string
		req  *{{.ProjectKebab}}.Request
		resp *{{.ProjectKebab}}.Response
		code codes.Code
	}{
		{
			name: "do the thing",
			req: &{{.ProjectKebab}}.Request{},
			resp: &{{.ProjectKebab}}.Response{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cc, err := grpc.NewClient("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
			require.NoError(t, err)
			defer func() {
				_ = cc.Close()
			}()

			cli := {{.ProjectKebab}}.NewAAAServiceClient(cc)

			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			resp, err := cli.Do(ctx, tt.req)
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
