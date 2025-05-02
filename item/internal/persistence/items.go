package persistence

import (
	"context"
	"errors"
	"time"

	"github.com/nats-io/nats.go/jetstream"
	"github.com/purplepudding/bricks/item/internal/core/catalog"
	"github.com/purplepudding/bricks/item/internal/core/model"
	persistencev1 "github.com/purplepudding/bricks/item/pkg/pb/bricks/item/v1/persistence"
	"github.com/purplepudding/bricks/lib/clients/natscli"
	"github.com/purplepudding/bricks/lib/common"
	"google.golang.org/protobuf/proto"
)

var _ catalog.ItemPersistence = (*Items)(nil)

const (
	ItemBucket = "items"
)

type Items struct {
	js jetstream.JetStream
}

func New(cfg natscli.Config) (*Items, error) {
	js, err := natscli.NewJetStream(cfg)
	if err != nil {
		return nil, err
	}

	return &Items{js: js}, nil
}

func (i *Items) GetByID(ctx context.Context, id string) (*model.Item, error) {
	kv, err := i.js.KeyValue(ctx, ItemBucket)
	if err != nil {
		return nil, err
	}

	return doGet(ctx, id, kv)
}

// ListAll obtains all items from storage, using the supplied pagination information
func (i *Items) ListAll(ctx context.Context, page common.Page) ([]*model.Item, error) {
	//TODO there may be resuable elements in this that could be reused for ListAvailable - consider :)

	kv, err := i.js.KeyValue(ctx, ItemBucket)
	if err != nil {
		return nil, err
	}

	kl, err := kv.ListKeys(ctx)
	if err != nil {
		return nil, err
	}

	// If LastID set on pagination, skip until we get to that ID
	skip := page.LastID != ""

	var count uint32
	var items []*model.Item
	for key := range kl.Keys() {
		if skip { // Only perform this work while we skip through the list

			// Naively assuming keys will be sorted the same way on multiple calls - chances are this assumption doesn't
			// hold, and we'll need a better way of paging through this list (or a better source for this data)
			if page.LastID == key {
				skip = false
			}
			continue
		}

		item, err := doGet(ctx, key, nil)
		if err != nil {
			return nil, err
		}

		items = append(items, item)

		count++
		if count >= page.Count {
			err := kl.Stop()
			if err != nil {
				return nil, err
			}
		}
	}

	return items, nil
}

// ListAvailableAt currently isn't implemented - more complicated implementation, will need to be able to query on the
// availability schedule which will likely want a different storage structure for this data.
func (i *Items) ListAvailableAt(ctx context.Context, date time.Time, page common.Page) ([]*model.Item, error) {
	return nil, errors.New("unimplemented")
}

func doGet(ctx context.Context, id string, kv jetstream.KeyValue) (*model.Item, error) {
	kve, err := kv.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	s := new(persistencev1.Item)
	if err := proto.Unmarshal(kve.Value(), s); err != nil {
		return nil, err
	}

	return model.ItemFromPersistencePB(s, id, kve.Revision()), nil
}

func (i *Items) Update(ctx context.Context, item *model.Item) (*model.Item, error) {
	kv, err := i.js.KeyValue(ctx, ItemBucket)
	if err != nil {
		return nil, err
	}

	pb, err := proto.Marshal(item.IntoPersistencePB())
	if err != nil {
		return nil, err
	}

	newVer, err := kv.Update(ctx, item.ID, pb, item.Version)
	if err != nil {
		return nil, err
	}

	item.Version = newVer
	return item, nil
}
