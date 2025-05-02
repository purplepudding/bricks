package catalog

import (
	"context"
	"time"

	"github.com/purplepudding/bricks/item/internal/core/model"
	"github.com/purplepudding/bricks/lib/common"
	"google.golang.org/protobuf/types/known/structpb"
)

//go:generate go tool go.uber.org/mock/mockgen -source interfaces.go -destination mock/interfaces.go -package mock_catalog

type ItemPersistence interface {
	GetByID(ctx context.Context, id string) (*model.Item, error)
	ListAll(ctx context.Context, page common.Page) ([]*model.Item, error)
	ListAvailableAt(ctx context.Context, date time.Time, page common.Page) ([]*model.Item, error)
	Update(ctx context.Context, item *model.Item) (*model.Item, error)
}

type AssetBundleClient interface {
	AssetBundleForItem(ctx context.Context, itemID string) (map[string]*structpb.Value, error)
}

type ItemSettingsClient interface {
	SettingsForItem(ctx context.Context, itemID string) (map[string]*structpb.Value, error)
}
