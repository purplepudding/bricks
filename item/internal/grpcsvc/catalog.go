package grpcsvc

import (
	"context"
	"time"

	itemv1 "github.com/purplepudding/bricks/api/pkg/pb/bricks/v1/item"
	"github.com/purplepudding/bricks/item/internal/model"
)

var _ itemv1.CatalogServiceServer = (*CatalogService)(nil)

type CatalogService struct {
	itemv1.UnimplementedCatalogServiceServer

	logic CatalogLogic
}

func NewCatalogService(logic CatalogLogic) *CatalogService {
	return &CatalogService{
		logic: logic,
	}
}

func (c *CatalogService) Get(ctx context.Context, req *itemv1.GetRequest) (*itemv1.GetResponse, error) {
	item, err := c.logic.Get(ctx, req.Id)
	if err != nil {
		//TODO handle
		return nil, err
	}

	resp := &itemv1.GetResponse{
		Item:        item.IntoPB(),
		AssetBundle: item.Assets.IntoPB(),
		Parameters:  item.Parameters.IntoPB(),
	}

	return resp, nil
}

func (c *CatalogService) List(ctx context.Context, req *itemv1.ListRequest) (*itemv1.ListResponse, error) {
	items, err := c.logic.List(ctx)
	if err != nil {
		//TODO handle
		return nil, err
	}

	var respItems []*itemv1.Item
	for _, item := range items {
		respItems = append(respItems, item.IntoPB())
	}

	return &itemv1.ListResponse{Items: respItems}, nil
}

func (c *CatalogService) ListAvailable(ctx context.Context, req *itemv1.ListAvailableRequest) (*itemv1.ListAvailableResponse, error) {
	var timestampOverride *time.Time
	if req.RequestTimestamp != nil {
		t := req.RequestTimestamp.AsTime()
		timestampOverride = &t
	}

	items, err := c.logic.ListAvailable(ctx, timestampOverride)
	if err != nil {
		//TODO handle
		return nil, err
	}

	var respItems []*itemv1.Item
	for _, item := range items {
		respItems = append(respItems, item.IntoPB())
	}

	return &itemv1.ListAvailableResponse{Items: respItems}, nil
}

func (c *CatalogService) UpdateItem(ctx context.Context, req *itemv1.UpdateItemRequest) (*itemv1.UpdateItemResponse, error) {
	nu, err := c.logic.UpdateItem(ctx, model.ItemFromPB(req.Item))
	if err != nil {
		//TODO handle
		return nil, err
	}

	return &itemv1.UpdateItemResponse{Version: nu.Version}, nil
}
