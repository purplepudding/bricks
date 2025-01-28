package config

import (
	_ "embed"
)

//go:embed defaults.yaml
var DefaultCfg []byte

type Config struct {
}
