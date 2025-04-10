package config

import (
	_ "embed"

	"github.com/purplepudding/foundation/lib/clients/natscli"
	"github.com/purplepudding/foundation/lib/clients/valkeycli"
)

//go:embed defaults.yaml
var DefaultCfg []byte

type Config struct {
	NATS   natscli.Config
	Valkey valkeycli.Config
}
