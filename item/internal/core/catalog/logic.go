package catalog

import (
	"context"
	"time"

	"github.com/purplepudding/bricks/item/internal/core/model"
	"github.com/purplepudding/bricks/item/internal/grpcsvc"
	"github.com/purplepudding/bricks/lib/common"
	"golang.org/x/sync/errgroup"
	"google.golang.org/protobuf/types/known/structpb"
)

var _ grpcsvc.CatalogLogic = (*Logic)(nil)

type Logic struct {
	persistence        ItemPersistence
	assetBundleClient  AssetBundleClient
	itemSettingsClient ItemSettingsClient
}

func New(p ItemPersistence, abc AssetBundleClient, isc ItemSettingsClient) *Logic {
	return &Logic{
		persistence:        p,
		assetBundleClient:  abc,
		itemSettingsClient: isc,
	}
}

func (l *Logic) Get(ctx context.Context, id string) (*model.Item, error) {
	var item *model.Item
	var assets, parameters map[string]*structpb.Value

	eg, ctx := errgroup.WithContext(ctx)

	// Pull the appropriate record from persistence
	eg.Go(func() (err error) {
		item, err = l.persistence.GetByID(ctx, id)
		return err
	})

	// Concurrently get the asset bundle set for this id from Assets
	eg.Go(func() (err error) {
		assets, err = l.assetBundleClient.AssetBundleForItem(ctx, id)
		return err
	})

	// Concurrently get the parameters set for this id from Settings
	eg.Go(func() (err error) {
		parameters, err = l.itemSettingsClient.SettingsForItem(ctx, id)
		return err
	})

	err := eg.Wait()
	if err != nil {
		return nil, err
	}

	item.Assets = assets
	item.Parameters = parameters
	return item, nil
}

func (l *Logic) List(ctx context.Context, page *common.Page) ([]*model.Item, error) {
	// Pull all records from persistence
	items, err := l.persistence.ListAll(ctx, pageOrDefault(page))
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (l *Logic) ListAvailable(ctx context.Context, timestampOverride *time.Time, page *common.Page) ([]*model.Item, error) {
	date := time.Now().UTC()

	if timestampOverride != nil {
		date = *timestampOverride
	}

	// Pull all records from persistence, filtering on their availability (may want a cache of this data)
	items, err := l.persistence.ListAvailableAt(ctx, date, pageOrDefault(page))
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (l *Logic) UpdateItem(ctx context.Context, item *model.Item) (*model.Item, error) {
	updatedItem, err := l.persistence.Update(ctx, item)
	if err != nil {
		return nil, err
	}

	return updatedItem, nil
}

func pageOrDefault(page *common.Page) common.Page {
	if page != nil {
		return *page
	}

	return common.Page{Count: 50}
}
