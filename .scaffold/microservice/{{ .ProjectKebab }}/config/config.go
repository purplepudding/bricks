package config

import (
	_ "embed"

	"github.com/purplepudding/foundation/lib/config"
)

//go:embed defaults.yaml
var DefaultCfg []byte

type Config struct {
	config.Microservice
}
