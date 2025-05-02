package grpcsvc

import (
	"context"
	"time"

	"github.com/purplepudding/bricks/item/internal/core/model"
	"github.com/purplepudding/bricks/lib/common"
)

//go:generate go tool go.uber.org/mock/mockgen -source interfaces.go -destination mock/interfaces.go -package mock_grpcsvc

type CatalogLogic interface {
	Get(ctx context.Context, id string) (*model.Item, error)
	List(ctx context.Context, page *common.Page) ([]*model.Item, error)
	ListAvailable(ctx context.Context, timestampOverride *time.Time, page *common.Page) ([]*model.Item, error)
	UpdateItem(ctx context.Context, item *model.Item) (*model.Item, error)
}
