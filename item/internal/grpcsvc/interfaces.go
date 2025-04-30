package grpcsvc

import (
	"context"
	"time"

	"github.com/purplepudding/bricks/item/internal/model"
)

//go:generate go tool go.uber.org/mock/mockgen -source interfaces.go -destination mock/interfaces.go -package mock_grpcsvc

type CatalogLogic interface {
	Get(ctx context.Context, id string) (*model.Item, error)
	List(ctx context.Context) ([]*model.Item, error)
	ListAvailable(ctx context.Context, timestampOverride *time.Time) ([]*model.Item, error)
	UpdateItem(ctx context.Context, item *model.Item) (*model.Item, error)
}
