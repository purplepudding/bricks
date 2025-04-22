package config

import (
	_ "embed"

	authConfig "github.com/purplepudding/foundation/auth/config"
	persistenceConfig "github.com/purplepudding/foundation/persistence/config"
	settingsConfig "github.com/purplepudding/foundation/settings/config"
)

//go:embed defaults.yaml
var DefaultCfg []byte

type Config struct {
	Auth        authConfig.Config
	Persistence persistenceConfig.Config
	Settings    settingsConfig.Config
}
