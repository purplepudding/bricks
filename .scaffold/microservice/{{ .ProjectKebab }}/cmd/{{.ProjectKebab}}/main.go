package main

import (
	"github.com/purplepudding/foundation"
	"github.com/purplepudding/foundation/lib/microservice"
	"github.com/purplepudding/foundation/{{.ProjectKebab}}/internal/config"
	"github.com/purplepudding/foundation/{{.ProjectKebab}}/internal/service"
)

func main() {
	microservice.Launch("{{.ProjectKebab}}", foundation.Version, config.DefaultCfg, new(config.Config), new(service.Service))
}
