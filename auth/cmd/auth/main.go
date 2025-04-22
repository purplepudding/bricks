package main

import (
	"github.com/purplepudding/bricks"
	"github.com/purplepudding/bricks/auth/config"
	"github.com/purplepudding/bricks/auth/service"
	"github.com/purplepudding/bricks/lib/microservice"
)

func main() {
	microservice.Launch("auth", bricks.Version, config.DefaultCfg, new(config.Config), new(service.Service))
}
