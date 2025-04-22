package main

import (
	"github.com/purplepudding/bricks"
	"github.com/purplepudding/bricks/lib/microservice"
	"github.com/purplepudding/bricks/settings/config"
	"github.com/purplepudding/bricks/settings/service"
)

func main() {
	microservice.Launch("settings", bricks.Version, config.DefaultCfg, new(config.Config), new(service.Service))
}
