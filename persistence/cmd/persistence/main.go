package main

import (
	"github.com/purplepudding/foundation"
	"github.com/purplepudding/foundation/lib/microservice"
	"github.com/purplepudding/foundation/persistence/config"
	"github.com/purplepudding/foundation/persistence/service"
)

func main() {
	microservice.Launch("persistence", foundation.Version, config.DefaultCfg, new(config.Config), new(service.Service))
}
