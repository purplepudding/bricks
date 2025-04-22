package settings

import (
	"context"

	v1 "github.com/purplepudding/bricks/settings/internal/grpcsvc/v1"
	"google.golang.org/protobuf/types/known/structpb"
)

var _ v1.ServiceSettingsLogic = (*ServiceSettingsLogic)(nil)

type ServiceSettingsLogic struct {
	globalSettings GlobalSettings
	store          ServiceSettingsStore
}

func NewServiceSettingsLogic(globalSettings GlobalSettings, store ServiceSettingsStore) *ServiceSettingsLogic {
	return &ServiceSettingsLogic{store: store, globalSettings: globalSettings}
}

func (logic *ServiceSettingsLogic) GetForService(ctx context.Context, service string) (map[string]*structpb.Value, error) {
	globals, err := logic.globalSettings.Get(ctx)
	if err != nil {
		//TODO sentinels and wrapping
		return nil, err
	}

	settings, err := logic.store.Get(ctx, service)
	if err != nil {
		//TODO sentinels and wrapping
		return nil, err
	}

	// Override globals with service-specific settings
	for k, v := range settings {
		globals[k] = v
	}

	return globals, nil
}

func (logic *ServiceSettingsLogic) SetForService(ctx context.Context, service string, settings map[string]*structpb.Value) error {
	return logic.store.Set(ctx, service, settings)
}

type GlobalSettings interface {
	Get(ctx context.Context) (map[string]*structpb.Value, error)
}

type ServiceSettingsStore interface {
	Get(ctx context.Context, collection string) (map[string]*structpb.Value, error)
	Set(ctx context.Context, collection string, settings map[string]*structpb.Value) error
}
