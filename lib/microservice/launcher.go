package microservice

import (
	"context"
	"log/slog"

	"github.com/purplepudding/bricks/lib/config"
	"github.com/purplepudding/bricks/lib/observability"
	"go.opentelemetry.io/otel"
)

// TODO handle exit signals
func Launch[T config.Observable](name, version string, defaultCfg []byte, cfg T, svc Service[T]) {
	if err := config.Load(name, defaultCfg, cfg); err != nil {
		slog.Error("failed to load config", "error", err)
		panic("config loading failed")
	}

	if cfg.EnableObservability() {
		slog.Info("enabling observability export")

		tp, mp, err := observability.EnableExport(name, version)
		if err != nil {
			slog.Error("failed to enable observability exporter", "error", err)
			panic("observability setup failed")
		}

		defer func() {
			_ = tp.Shutdown(context.Background())
			_ = mp.Shutdown(context.Background())
		}()

		otel.SetTracerProvider(tp)
		otel.SetMeterProvider(mp)
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
