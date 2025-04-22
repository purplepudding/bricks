package storage

import (
	"context"

	"github.com/purplepudding/bricks/persistence/internal/core/model"
	"github.com/purplepudding/bricks/persistence/internal/grpcsvc"
	"google.golang.org/protobuf/types/known/structpb"
)

var _ grpcsvc.StorageLogic = (*Logic)(nil)

type Logic struct {
	persistence Persistence
}

func NewLogic(storage Persistence) *Logic {
	return &Logic{persistence: storage}
}

func (s *Logic) Get(ctx context.Context, key model.StorageKey) (*structpb.Struct, error) {
	//TODO map storage engine errors into service errors
	return s.persistence.Get(ctx, key)
}

func (s *Logic) Set(ctx context.Context, key model.StorageKey, value *structpb.Struct) error {
	//TODO map storage engine errors into service errors
	return s.persistence.Set(ctx, key, value)
}

type Persistence interface {
	Get(ctx context.Context, key model.StorageKey) (*structpb.Struct, error)
	Set(ctx context.Context, key model.StorageKey, value *structpb.Struct) error
}
