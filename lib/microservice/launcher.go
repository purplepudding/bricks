package microservice

import (
	"log/slog"

	"github.com/purplepudding/foundation/lib/config"
)

// TODO handle exit signals
func Launch(name string, cfg any, svc Service) {
	if err := config.Load(cfg); err != nil {
		slog.Error("failed to load config", "error", err)
		panic("config loading failed")
	}

	slog.Info("booting service", "name", name) //TODO add version number

	if err := svc.Wire(); err != nil {
		slog.Error("failed to wire the service", "error", err) //TODO replace with unified tracing/logging
		panic("startup failed")
	}

	slog.Info("running service", "name", name)

	if err := svc.Run(); err != nil {
		slog.Error("error while running the service", "error", err)
		panic("service runtime error")
	}
}
