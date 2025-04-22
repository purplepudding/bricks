package config

import (
	_ "embed"

	authConfig "github.com/purplepudding/bricks/auth/config"
	persistenceConfig "github.com/purplepudding/bricks/persistence/config"
	settingsConfig "github.com/purplepudding/bricks/settings/config"
)

//go:embed defaults.yaml
var DefaultCfg []byte

type Config struct {
	Auth        authConfig.Config
	Persistence persistenceConfig.Config
	Settings    settingsConfig.Config
}
