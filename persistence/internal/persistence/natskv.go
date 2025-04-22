package persistence

import (
	"context"
	"errors"
	"fmt"

	"github.com/nats-io/nats.go/jetstream"
	"github.com/purplepudding/bricks/lib/clients/natscli"
	"github.com/purplepudding/bricks/persistence/internal/core/model"
	"github.com/purplepudding/bricks/persistence/internal/core/storage"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/structpb"
)

var _ storage.Persistence = (*NatsKVPersistence)(nil)

type NatsKVPersistence struct {
	js jetstream.JetStream
}

func NewNatsKVPersistence(cfg natscli.Config) (*NatsKVPersistence, error) {
	js, err := natscli.NewJetStream(cfg)
	if err != nil {
		return nil, err
	}

	return &NatsKVPersistence{js: js}, nil
}

func (n *NatsKVPersistence) Get(ctx context.Context, key model.StorageKey) (*structpb.Struct, error) {
	kv, err := n.js.KeyValue(ctx, key.Category())
	if err != nil {
		return nil, err
	}

	kve, err := kv.Get(ctx, key.Key())
	if err != nil {
		return nil, err
	}

	s := new(structpb.Struct)
	if err := proto.Unmarshal(kve.Value(), s); err != nil {
		return nil, err
	}

	return s, nil
}

func (n *NatsKVPersistence) Set(ctx context.Context, key model.StorageKey, value *structpb.Struct) error {
	b, err := proto.Marshal(value)
	if err != nil {
		return err
	}

	kv, err := n.js.KeyValue(ctx, key.Category())
	if err != nil {
		if !errors.Is(err, jetstream.ErrBucketNotFound) {
			return err
		}

		// As the bucket doesn't exist, we need to create it first and then continue
		kv, err = n.js.CreateKeyValue(ctx, jetstream.KeyValueConfig{
			Bucket: key.Category(),
		})
		if err != nil {
			return fmt.Errorf("failed to create bucket %q: %w", key.Category(), err)
		}
	}

	if _, err := kv.Put(ctx, key.Key(), b); err != nil {
		return err
	}

	return nil
}
