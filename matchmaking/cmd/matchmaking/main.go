package main

import (
	"github.com/purplepudding/bricks"
	"github.com/purplepudding/bricks/lib/microservice"
	"github.com/purplepudding/bricks/matchmaking/config"
	"github.com/purplepudding/bricks/matchmaking/service"
)

func main() {
	microservice.Launch("matchmaking", bricks.Version, config.DefaultCfg, new(config.Config), new(service.Service))
}
