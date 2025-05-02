package assetbundles

import (
	"context"

	"github.com/purplepudding/bricks/item/internal/core/catalog"
	"google.golang.org/protobuf/types/known/structpb"
)

var _ catalog.AssetBundleClient = (*Client)(nil)

type Client struct{}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) AssetBundleForItem(ctx context.Context, itemID string) (map[string]*structpb.Value, error) {
	//TODO wire to asset bundle service once implemented
	return map[string]*structpb.Value{}, nil
}
