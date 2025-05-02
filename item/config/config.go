package config

import (
	_ "embed"

	"github.com/purplepudding/bricks/lib/clients/natscli"
	"github.com/purplepudding/bricks/lib/clients/settingscli"
	"github.com/purplepudding/bricks/lib/config"
)

//go:embed defaults.yaml
var DefaultCfg []byte

type Config struct {
	config.Microservice `koanf:",squash"`

	Clients Clients
}

type Clients struct {
	NATS     natscli.Config
	Settings settingscli.Config
}
