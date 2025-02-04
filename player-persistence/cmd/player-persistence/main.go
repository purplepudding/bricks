package main

import (
	"github.com/purplepudding/foundation"
	"github.com/purplepudding/foundation/lib/microservice"
	"github.com/purplepudding/foundation/player-persistence/internal/config"
	"github.com/purplepudding/foundation/player-persistence/internal/service"
)

func main() {
	microservice.Launch("player-persistence", foundation.Version, config.DefaultCfg, new(config.Config), new(service.Service))
}
