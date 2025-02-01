package settings

import (
	"context"
	"errors"
	"strings"

	"github.com/knadh/koanf"
	"github.com/purplepudding/foundation/api/pkg/pb/foundation/v1/settings"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/structpb"
)

var _ koanf.Provider = (*Provider)(nil)

type Provider struct {
	service string
	cli     settings.ServiceSettingsServiceClient
}

func NewProvider(settingsSvcUrl, service string) koanf.Provider {
	cc, err := grpc.NewClient(settingsSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		// TODO confirm this is only ever a fatal error
		panic(err)
	}
	//TODO we might want a close method to close cc gracefully on shutdown

	cli := settings.NewServiceSettingsServiceClient(cc)

	return &Provider{
		cli:     cli,
		service: service,
	}
}

func (p *Provider) ReadBytes() ([]byte, error) {
	return nil, errors.New("not supported")
}

func (p *Provider) Read() (map[string]interface{}, error) {
	resp, err := p.cli.GetServiceSettings(context.Background(), &settings.GetServiceSettingsRequest{
		Service: p.service,
	})
	if err != nil {
		return nil, err
	}

	return p.expand(resp.GetSettings()), nil
}

func (p *Provider) expand(settings map[string]*structpb.Value) map[string]interface{} {
	res := make(map[string]any)
	for k, v := range settings {
		var val any
		// Read the string backwards, appending the value to the end
		split := strings.Split(k, ":")
		for i := len(split) - 1; i > 0; i-- {
			if val == nil {
				val = map[string]any{split[i]: v.AsInterface()}
				continue
			}

			val = map[string]any{split[i]: val}

		}

		if val == nil {
			val = v.AsInterface()
		}
		res[split[0]] = val
	}

	return res
}
