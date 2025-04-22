package main

import (
	"github.com/purplepudding/foundation"
	"github.com/purplepudding/foundation/lib/microservice"
	"github.com/purplepudding/foundation/monolith/config"
	"github.com/purplepudding/foundation/monolith/service"
)

func main() {
	microservice.Launch("monolith", foundation.Version, config.DefaultCfg, new(config.Config), new(service.Service))
}
