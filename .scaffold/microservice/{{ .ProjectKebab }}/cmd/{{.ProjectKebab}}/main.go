package main

import (
	"github.com/purplepudding/bricks"
	"github.com/purplepudding/bricks/lib/microservice"
	"github.com/purplepudding/bricks/{{.ProjectKebab}}/internal/config"
	"github.com/purplepudding/bricks/{{.ProjectKebab}}/internal/service"
)

func main() {
	microservice.Launch("{{.ProjectKebab}}", bricks.Version, config.DefaultCfg, new(config.Config), new(service.Service))
}
