package persistence

import (
	"context"

	"github.com/purplepudding/foundation/persistence/internal/core/model"
	"github.com/purplepudding/foundation/persistence/internal/core/storage"
	"github.com/valkey-io/valkey-go"
	"github.com/vmihailenco/msgpack/v5"
	"google.golang.org/protobuf/types/known/structpb"
)

var _ storage.Persistence = (*ValkeyStorage)(nil)

type ValkeyStorage struct {
	valkeyCli valkey.Client
}

func NewValkeyStorage(valkeyCli valkey.Client) *ValkeyStorage {
	return &ValkeyStorage{valkeyCli: valkeyCli}
}

func (s *ValkeyStorage) Get(ctx context.Context, key model.StorageKey) (*structpb.Struct, error) {
	res := s.valkeyCli.Do(ctx, s.valkeyCli.B().Hget().Key(key.Key()).Field(key.Category()).Build())

	b, err := res.AsBytes()
	if err != nil {
		//TODO sentinel and wrapping
		return nil, err
	}

	var m map[string]any
	err = msgpack.Unmarshal(b, &m)
	if err != nil {
		//TODO sentinel and wrapping
		return nil, err
	}

	sv, err := structpb.NewStruct(m)
	if err != nil {
		//TODO sentinel and wrapping
		return nil, err
	}

	return sv, nil
}

func (s *ValkeyStorage) Set(ctx context.Context, key model.StorageKey, value *structpb.Struct) error {
	b, err := msgpack.Marshal(value.AsMap())
	if err != nil {
		return err
	}

	res := s.valkeyCli.Do(ctx, s.valkeyCli.B().Hset().Key(key.Key()).FieldValue().FieldValue(key.Category(), string(b)).Build())

	return res.Error()
}
