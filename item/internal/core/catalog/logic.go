package catalog

import (
	"context"
	"time"

	"github.com/purplepudding/bricks/item/internal/core/model"
	"github.com/purplepudding/bricks/item/internal/grpcsvc"
)

var _ grpcsvc.CatalogLogic = (*Logic)(nil)

type Logic struct{}

func New() *Logic {
	return &Logic{}
}

func (l *Logic) Get(ctx context.Context, id string) (*model.Item, error) {
	// Pull the appropriate record from persistence
	// Concurrently get the asset bundle set for this id from Assets
	// Concurrently get the parameters set for this id from Settings
	
	//TODO implement me
	panic("implement me")
}

func (l *Logic) List(ctx context.Context) ([]*model.Item, error) {
	// Pull all records from persistence

	//TODO implement me
	panic("implement me")
}

func (l *Logic) ListAvailable(ctx context.Context, timestampOverride *time.Time) ([]*model.Item, error) {
	// Pull all records from persistence, filtering on their availability (may want a cache of this data)

	//TODO implement me
	panic("implement me")
}

func (l *Logic) UpdateItem(ctx context.Context, item *model.Item) (*model.Item, error) {
	//TODO implement me
	panic("implement me")
}
