package main

import (
	"github.com/purplepudding/bricks"
	"github.com/purplepudding/bricks/lib/microservice"
	"github.com/purplepudding/bricks/persistence/config"
	"github.com/purplepudding/bricks/persistence/service"
)

func main() {
	microservice.Launch("persistence", bricks.Version, config.DefaultCfg, new(config.Config), new(service.Service))
}
