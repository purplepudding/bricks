package microservice

import (
	"log/slog"
)

// TODO handle exit signals
func Launch(name string, svc Service) {
	slog.Info("booting service", "name", name) //TODO add version from config

	if err := svc.Wire(); err != nil {
		slog.Error("failed to wire the service", "error", err) //TODO replace with unified tracing/logging
		panic("startup failed")
	}

	if err := svc.Run(); err != nil {
		slog.Error("error while running the service", "error", err)
		panic("service runtime error")
	}
}
