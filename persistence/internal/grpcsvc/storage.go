package grpcsvc

import (
	"context"
	"errors"

	persistencev1 "github.com/purplepudding/bricks/api/pkg/pb/bricks/v1/persistence"
	"github.com/purplepudding/bricks/persistence/internal/core/model"
	"google.golang.org/protobuf/types/known/structpb"
)

var _ persistencev1.StorageServiceServer = (*StorageService)(nil)

type StorageService struct {
	persistencev1.UnimplementedStorageServiceServer

	logic StorageLogic
}

func NewStorageService(logic StorageLogic) *StorageService {
	return &StorageService{logic: logic}
}

func (s *StorageService) Set(ctx context.Context, req *persistencev1.SetRequest) (*persistencev1.SetResponse, error) {
	key, err := intoModelKey(req.Key)
	if err != nil {
		//TODO map to grpc error type
		return nil, err
	}

	if err := s.logic.Set(ctx, key, req.Value); err != nil {
		//TODO map to correct error type
		return nil, err
	}

	return &persistencev1.SetResponse{}, nil
}

func (s *StorageService) Get(ctx context.Context, req *persistencev1.GetRequest) (*persistencev1.GetResponse, error) {
	key, err := intoModelKey(req.Key)
	if err != nil {
		//TODO map to grpc error type
		return nil, err
	}

	res, err := s.logic.Get(ctx, key)
	if err != nil {
		//TODO map to correct error types
		return nil, err
	}

	return &persistencev1.GetResponse{Value: res}, err
}

func intoModelKey(reqKey *persistencev1.Key) (model.StorageKey, error) {
	if pkey := reqKey.GetPlayerKey(); pkey != nil {
		return model.PlayerStorageKey{
			TitleID:  pkey.TitleId,
			PlayerID: pkey.PlayerId,
			Datatype: pkey.Datatype,
		}, nil
	}

	if tkey := reqKey.GetTitleKey(); tkey != nil {
		return model.TitleStorageKey{
			TitleID:  tkey.TitleId,
			Datatype: tkey.Datatype,
		}, nil
	}

	//TODO use a sentinel error and clean up the typing
	return nil, errors.New("key type unhandled")
}

type StorageLogic interface {
	Get(ctx context.Context, key model.StorageKey) (*structpb.Struct, error)
	Set(ctx context.Context, key model.StorageKey, value *structpb.Struct) error
}
