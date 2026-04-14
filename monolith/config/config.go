package config

import (
	_ "embed"

	authConfig "github.com/purplepudding/bricks/auth/config"
	itemConfig "github.com/purplepudding/bricks/item/config"
	"github.com/purplepudding/bricks/lib/config"
	matchmakingConfig "github.com/purplepudding/bricks/matchmaking/config"
	persistenceConfig "github.com/purplepudding/bricks/persistence/config"
	settingsConfig "github.com/purplepudding/bricks/settings/config"
)

//go:embed defaults.yaml
var DefaultCfg []byte

type Config struct {
	config.Microservice `koanf:",squash"`
	Auth                authConfig.Config
	Item                itemConfig.Config
	Matchmaking         matchmakingConfig.Config
	Persistence         persistenceConfig.Config
	Settings            settingsConfig.Config
	Gateway             config.Microservice
}
