package config

import (
	_ "embed"

	"github.com/purplepudding/foundation/lib/valkeycli"
)

//go:embed defaults.yaml
var DefaultCfg []byte

type Config struct {
	Valkey valkeycli.Config
}
