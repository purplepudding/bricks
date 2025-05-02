package settings

import (
	"context"
	"errors"
	"strings"

	"github.com/knadh/koanf"
	"github.com/purplepudding/bricks/api/pkg/pb/bricks/v1/settings"
	"github.com/purplepudding/bricks/lib/clients/settingscli"
	"google.golang.org/protobuf/types/known/structpb"
)

var _ koanf.Provider = (*Provider)(nil)

type Provider struct {
	service string
	cli     *settingscli.Client
}

func NewProvider(settingsSvcUrl, service string) koanf.Provider {
	cli, err := settingscli.New(settingscli.Config{Addr: settingsSvcUrl})
	if err != nil {
		// TODO confirm this is only ever a fatal error
		panic(err)
	}

	return &Provider{
		cli:     cli,
		service: service,
	}
}

func (p *Provider) ReadBytes() ([]byte, error) {
	return nil, errors.New("not supported")
}

func (p *Provider) Read() (map[string]interface{}, error) {
	resp, err := p.cli.ServiceSettingsClient.GetServiceSettings(context.Background(), &settings.GetServiceSettingsRequest{
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
