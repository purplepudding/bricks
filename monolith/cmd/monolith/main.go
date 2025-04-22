package main

import (
	"github.com/purplepudding/bricks"
	"github.com/purplepudding/bricks/lib/microservice"
	"github.com/purplepudding/bricks/monolith/config"
	"github.com/purplepudding/bricks/monolith/service"
)

func main() {
	microservice.Launch("monolith", bricks.Version, config.DefaultCfg, new(config.Config), new(service.Service))
}
