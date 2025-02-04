package grpcsvc

import (
	"context"

	playerpersistencev1 "github.com/purplepudding/foundation/api/pkg/pb/foundation/v1/playerpersistence"
	"google.golang.org/protobuf/types/known/structpb"
)

var _ playerpersistencev1.StorageServiceServer = (*StorageService)(nil)

type StorageService struct {
	playerpersistencev1.UnimplementedStorageServiceServer

	logic StorageLogic
}

func NewStorageService(logic StorageLogic) *StorageService {
	return &StorageService{logic: logic}
}

func (s *StorageService) Set(ctx context.Context, req *playerpersistencev1.SetRequest) (*playerpersistencev1.SetResponse, error) {
	if err := s.logic.Set(ctx, req.PlayerId, req.Datatype, req.Value); err != nil {
		//TODO map to correct error type
		return nil, err
	}

	return &playerpersistencev1.SetResponse{}, nil
}

func (s *StorageService) Get(ctx context.Context, req *playerpersistencev1.GetRequest) (*playerpersistencev1.GetResponse, error) {
	res, err := s.logic.Get(ctx, req.PlayerId, req.Datatype)
	if err != nil {
		//TODO map to correct error types
		return nil, err
	}

	return &playerpersistencev1.GetResponse{Value: res}, err
}

type StorageLogic interface {
	Get(ctx context.Context, playerID string, datatype string) (*structpb.Struct, error)
	Set(ctx context.Context, playerID string, datatype string, value *structpb.Struct) error
}
