package main

import (
	"github.com/purplepudding/foundation"
	"github.com/purplepudding/foundation/auth/config"
	"github.com/purplepudding/foundation/auth/service"
	"github.com/purplepudding/foundation/lib/microservice"
)

func main() {
	microservice.Launch("auth", foundation.Version, config.DefaultCfg, new(config.Config), new(service.Service))
}
