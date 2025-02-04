package storage

import (
	"context"

	"github.com/purplepudding/foundation/player-persistence/internal/grpcsvc"
	"google.golang.org/protobuf/types/known/structpb"
)

var _ grpcsvc.StorageLogic = (*Logic)(nil)

type Logic struct {
	persistence Persistence
}

func NewLogic(storage Persistence) *Logic {
	return &Logic{persistence: storage}
}

//TODO add event emission here. may want to consider roundtripping the data to make events the source of truth.

func (s *Logic) Get(ctx context.Context, playerID string, datatype string) (*structpb.Struct, error) {
	return s.persistence.Get(ctx, playerID, datatype)
}

func (s *Logic) Set(ctx context.Context, playerID string, datatype string, value *structpb.Struct) error {
	return s.persistence.Set(ctx, playerID, datatype, value)
}

type Persistence interface {
	Get(ctx context.Context, playerID string, datatype string) (*structpb.Struct, error)
	Set(ctx context.Context, playerID string, datatype string, value *structpb.Struct) error
}
