package settings

import (
	"context"

	settingsv1 "github.com/purplepudding/foundation/api/pkg/pb/foundation/v1/settings"
	v1 "github.com/purplepudding/foundation/settings/internal/grpcsvc/v1"
)

const (
	GlobalCollection = "global"
)

var _ v1.GlobalSettingsLogic = (*GlobalSettingsLogic)(nil)

type GlobalSettingsLogic struct {
	store SettingsStore
}

func NewSettingsLogic(store SettingsStore) *GlobalSettingsLogic {
	return &GlobalSettingsLogic{store: store}
}

func (logic *GlobalSettingsLogic) SetSettings(ctx context.Context, settings []*settingsv1.Settings) error {
	entries := make(map[string]any)
	for _, setting := range settings {
		entries[setting.Key] = setting.Value.AsInterface()
	}

	return logic.store.Set(ctx, GlobalCollection, entries)
}

type SettingsStore interface {
	Set(ctx context.Context, collection string, entries map[string]any) error
}
