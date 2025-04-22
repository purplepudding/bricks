package settings

import (
	"context"
	"errors"
	"testing"

	"github.com/purplepudding/bricks/api/pkg/pb/bricks/v1/settings"
	mocksettings "github.com/purplepudding/bricks/api/pkg/pb/bricks/v1/settings/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"google.golang.org/protobuf/types/known/structpb"
)

func TestProvider_Read(t *testing.T) {
	var (
		ctx     = context.Background()
		service = "service"
		errTest = errors.New("test error")
	)

	tests := []struct {
		name   string
		setup  func(mockCli *mocksettings.MockServiceSettingsServiceClient)
		result map[string]any
		err    error
	}{
		{
			name: "when an error occurs, return it",
			setup: func(mockCli *mocksettings.MockServiceSettingsServiceClient) {
				mockCli.EXPECT().GetServiceSettings(ctx, gomock.Any()).Return(nil, errTest)
			},
			err: errTest,
		},
		{
			name: "when a single level map is returned, return it",
			setup: func(mockCli *mocksettings.MockServiceSettingsServiceClient) {
				mockCli.EXPECT().GetServiceSettings(ctx, gomock.Any()).Return(&settings.GetServiceSettingsResponse{
					Settings: map[string]*structpb.Value{
						"foo": structpb.NewStringValue("bar"),
					},
				}, nil)
			},
			result: map[string]any{
				"foo": "bar",
			},
		},
		{
			name: "when a multi-level map is returned, return it",
			setup: func(mockCli *mocksettings.MockServiceSettingsServiceClient) {
				mockCli.EXPECT().GetServiceSettings(ctx, gomock.Any()).Return(&settings.GetServiceSettingsResponse{
					Settings: map[string]*structpb.Value{
						"baz:struct:foo": structpb.NewStringValue("bar"),
					},
				}, nil)
			},
			result: map[string]any{
				"baz": map[string]any{
					"struct": map[string]any{
						"foo": "bar",
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockCli := mocksettings.NewMockServiceSettingsServiceClient(ctrl)

			if tt.setup != nil {
				tt.setup(mockCli)
			}

			p := &Provider{service: service, cli: mockCli}

			res, err := p.Read()
			if tt.err == nil {
				assert.NoError(t, err)
				assert.Equal(t, tt.result, res)
			} else {
				assert.Nil(t, res)
				assert.Error(t, err)
				assert.ErrorIs(t, err, tt.err)
			}
		})
	}
}
