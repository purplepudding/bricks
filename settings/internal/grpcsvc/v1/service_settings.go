package v1

import (
	"context"

	settingsv1 "github.com/purplepudding/foundation/api/pkg/pb/foundation/v1/settings"
	"google.golang.org/protobuf/types/known/structpb"
)

var _ settingsv1.ServiceSettingsServiceServer = (*ServiceSettingsService)(nil)

type ServiceSettingsService struct {
	settingsv1.UnimplementedServiceSettingsServiceServer

	serviceSettings ServiceSettingsLogic
}

func NewServiceSettingsService(serviceSettings ServiceSettingsLogic) *ServiceSettingsService {
	return &ServiceSettingsService{
		serviceSettings: serviceSettings,
	}
}

func (g *ServiceSettingsService) GetServiceSettings(ctx context.Context, req *settingsv1.GetServiceSettingsRequest) (*settingsv1.GetServiceSettingsResponse, error) {
	settings, err := g.serviceSettings.GetForService(ctx, req.Service)
	if err != nil {
		//TODO handle proper error cases
		return nil, err
	}

	return &settingsv1.GetServiceSettingsResponse{
		Settings: settings,
	}, nil
}

type ServiceSettingsLogic interface {
	GetForService(ctx context.Context, service string) (map[string]*structpb.Value, error)
}
