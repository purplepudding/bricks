package persistence

import (
	"context"

	"github.com/nats-io/nats.go/jetstream"
	"github.com/purplepudding/foundation/lib/clients/natscli"
	"github.com/purplepudding/foundation/persistence/internal/core/model"
	"github.com/purplepudding/foundation/persistence/internal/core/storage"
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
		//TODO if kv bucket doesn't exist, create it and continue
		return err
	}

	if _, err := kv.Put(ctx, key.Key(), b); err != nil {
		return err
	}

	return nil
}
