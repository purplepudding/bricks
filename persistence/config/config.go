package config

import (
	_ "embed"

	"github.com/purplepudding/bricks/lib/clients/natscli"
	"github.com/purplepudding/bricks/lib/clients/valkeycli"
	"github.com/purplepudding/bricks/lib/config"
)

//go:embed defaults.yaml
var DefaultCfg []byte

type Config struct {
	config.Microservice `koanf:",squash"`

	NATS   natscli.Config
	Valkey valkeycli.Config
}
