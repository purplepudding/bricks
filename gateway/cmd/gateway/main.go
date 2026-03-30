package main

import (
	"github.com/purplepudding/bricks"
	"github.com/purplepudding/bricks/lib/microservice"
	"github.com/purplepudding/bricks/gateway/config"
	"github.com/purplepudding/bricks/gateway/service"
)

func main() {
	microservice.Launch("gateway", bricks.Version, config.DefaultCfg, new(config.Config), new(service.Service))
}
