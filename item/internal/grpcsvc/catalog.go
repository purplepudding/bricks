package grpcsvc

import (
	itemv1 "github.com/purplepudding/bricks/api/pkg/pb/bricks/v1/item"
)

var _ itemv1.CatalogServiceServer = (*CatalogService)(nil)

type CatalogService struct {
	itemv1.UnimplementedCatalogServiceServer
}
