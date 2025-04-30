package main

import (
	"github.com/purplepudding/bricks"
	"github.com/purplepudding/bricks/lib/microservice"
	"github.com/purplepudding/bricks/item/config"
	"github.com/purplepudding/bricks/item/service"
)

func main() {
	microservice.Launch("item", bricks.Version, config.DefaultCfg, new(config.Config), new(service.Service))
}
