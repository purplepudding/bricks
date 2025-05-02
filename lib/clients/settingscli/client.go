package settingscli

import (
	settingsv1 "github.com/purplepudding/bricks/api/pkg/pb/bricks/v1/settings"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func New(cfg Config) (*Client, error) {
	cc, err := grpc.NewClient(cfg.Addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
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
