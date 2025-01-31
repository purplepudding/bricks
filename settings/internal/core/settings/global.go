package settings

import (
	"context"

	v1 "github.com/purplepudding/foundation/settings/internal/grpcsvc/v1"
	"google.golang.org/protobuf/types/known/structpb"
)

const (
	GlobalCollection = "global"
)

var _ v1.GlobalSettingsLogic = (*GlobalSettingsLogic)(nil)
var _ GlobalSettings = (*GlobalSettingsLogic)(nil)

type GlobalSettingsLogic struct {
	store GlobalSettingsStore
}

func NewGlobalSettingsLogic(store GlobalSettingsStore) *GlobalSettingsLogic {
	return &GlobalSettingsLogic{store: store}
}

func (logic *GlobalSettingsLogic) Get(ctx context.Context) (map[string]*structpb.Value, error) {
	return logic.store.Get(ctx, GlobalCollection)
}

func (logic *GlobalSettingsLogic) SetSettings(ctx context.Context, settings map[string]*structpb.Value) error {
	//TODO consider what logic we could bring up to this level, if any
	return logic.store.Set(ctx, GlobalCollection, settings)
}

type GlobalSettingsStore interface {
	Get(ctx context.Context, collection string) (map[string]*structpb.Value, error)
	Set(ctx context.Context, collection string, entries map[string]*structpb.Value) error
}
