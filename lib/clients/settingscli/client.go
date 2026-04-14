package settingscli

import (
	settingsv1 "github.com/purplepudding/bricks/api/pkg/pb/bricks/v1/settings"
	"github.com/purplepudding/bricks/lib/clients/grpccli"
)

func New(cfg Config) (*Client, error) {
	cc, err := grpccli.New(cfg.Addr)
	if err != nil {
		return nil, err
	}

	return &Client{
		GlobalSettingsClient:  settingsv1.NewGlobalSettingsServiceClient(cc),
		ServiceSettingsClient: settingsv1.NewServiceSettingsServiceClient(cc),
		ItemParamtersClient:   settingsv1.NewItemParametersServiceClient(cc),
	}, nil
}

type Client struct {
	GlobalSettingsClient  settingsv1.GlobalSettingsServiceClient
	ServiceSettingsClient settingsv1.ServiceSettingsServiceClient
	ItemParamtersClient   settingsv1.ItemParametersServiceClient
}
