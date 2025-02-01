package microservice

import (
	"log/slog"

	"github.com/purplepudding/foundation/lib/config"
)

// TODO handle exit signals
func Launch[T any](name, version string, defaultCfg []byte, cfg T, svc Service[T]) {
	if err := config.Load(name, defaultCfg, cfg); err != nil {
		slog.Error("failed to load config", "error", err)
		panic("config loading failed")
	}

	slog.Info("booting service", "name", name, "version", version)

	if err := svc.Wire(cfg); err != nil {
		slog.Error("failed to wire the service", "error", err) //TODO replace with unified tracing/logging
		panic("startup failed")
	}

	slog.Info("running service", "name", name, "version", version)

	if err := svc.Run(); err != nil {
		slog.Error("error while running the service", "error", err)
		panic("service runtime error")
	}
}
