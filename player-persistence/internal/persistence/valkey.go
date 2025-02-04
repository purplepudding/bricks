package persistence

import (
	"context"

	"github.com/purplepudding/foundation/player-persistence/internal/core/storage"
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

func (s *ValkeyStorage) Get(ctx context.Context, playerID string, datatype string) (*structpb.Struct, error) {
	res := s.valkeyCli.Do(ctx, s.valkeyCli.B().Hget().Key(playerID).Field(datatype).Build())

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

func (s *ValkeyStorage) Set(ctx context.Context, playerID string, datatype string, value *structpb.Struct) error {
	b, err := msgpack.Marshal(value.AsMap())
	if err != nil {
		return err
	}

	res := s.valkeyCli.Do(ctx, s.valkeyCli.B().Hset().Key(playerID).FieldValue().FieldValue(datatype, string(b)).Build())

	return res.Error()
}
