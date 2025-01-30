package grpcsvc

import (
	"context"

	settingsv1 "github.com/purplepudding/foundation/api/pkg/pb/foundation/v1/settings"
)

var _ settingsv1.GlobalSettingsServiceServer = (*GlobalSettingsService)(nil)

type GlobalSettingsService struct {
	settingsv1.UnimplementedGlobalSettingsServiceServer
}

func (g *GlobalSettingsService) SetGlobalSettings(ctx context.Context, req *settingsv1.SetGlobalSettingsRequest) (*settingsv1.SetGlobalSettingsResponse, error) {
	//TODO implement me
	panic("implement me")
}
