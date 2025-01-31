package v1

import (
	"context"

	settingsv1 "github.com/purplepudding/foundation/api/pkg/pb/foundation/v1/settings"
	"google.golang.org/protobuf/types/known/structpb"
)

var _ settingsv1.GlobalSettingsServiceServer = (*GlobalSettingsService)(nil)

type GlobalSettingsService struct {
	settingsv1.UnimplementedGlobalSettingsServiceServer

	globalSettings GlobalSettingsLogic
}

func NewGlobalSettingsService(globalSettings GlobalSettingsLogic) *GlobalSettingsService {
	return &GlobalSettingsService{
		globalSettings: globalSettings,
	}
}

func (g *GlobalSettingsService) SetGlobalSettings(ctx context.Context, req *settingsv1.SetGlobalSettingsRequest) (*settingsv1.SetGlobalSettingsResponse, error) {
	if err := g.globalSettings.SetSettings(ctx, req.Settings); err != nil {
		//TODO handle proper error cases
		return nil, err
	}

	return &settingsv1.SetGlobalSettingsResponse{}, nil
}

type GlobalSettingsLogic interface {
	SetSettings(ctx context.Context, settings map[string]*structpb.Value) error
}
