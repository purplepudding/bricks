package main

import (
	"github.com/purplepudding/foundation"
	"github.com/purplepudding/foundation/lib/microservice"
	"github.com/purplepudding/foundation/settings/internal/config"
	"github.com/purplepudding/foundation/settings/internal/service"
)

func main() {
	microservice.Launch("settings", foundation.Version, config.DefaultCfg, new(config.Config), new(service.Service))
}
