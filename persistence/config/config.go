package config

import (
	_ "embed"

	"github.com/purplepudding/foundation/lib/clients/natscli"
	"github.com/purplepudding/foundation/lib/clients/valkeycli"
	"github.com/purplepudding/foundation/lib/config"
)

//go:embed defaults.yaml
var DefaultCfg []byte

type Config struct {
	config.Microservice

	NATS   natscli.Config
	Valkey valkeycli.Config
}
