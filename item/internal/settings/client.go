package settings

import (
	"context"

	settingsv1 "github.com/purplepudding/bricks/api/pkg/pb/bricks/v1/settings"
	"github.com/purplepudding/bricks/item/internal/core/catalog"
	"github.com/purplepudding/bricks/lib/clients/settingscli"
	"google.golang.org/protobuf/types/known/structpb"
)

var _ catalog.ItemSettingsClient = (*Client)(nil)

type Client struct {
	cli *settingscli.Client
}

func NewClient(cfg settingscli.Config) *Client {
	cli, err := settingscli.New(cfg)
	if err != nil {
		// TODO confirm this is only ever a fatal error
		panic(err)
	}

	return &Client{
		cli: cli,
	}
}

func (c *Client) SettingsForItem(ctx context.Context, itemID string) (map[string]*structpb.Value, error) {
	params, err := c.cli.ItemParamtersClient.GetItemParameters(ctx, &settingsv1.GetItemParametersRequest{
		ItemId: itemID,
	})
	if err != nil {
		return nil, err
	}

	return params.Parameters, nil
}
