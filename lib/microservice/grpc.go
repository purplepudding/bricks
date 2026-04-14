package microservice

import (
	"github.com/purplepudding/bricks/lib/config"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"google.golang.org/grpc"
	oteltracing "google.golang.org/grpc/experimental/opentelemetry"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/stats/opentelemetry"
)

type GRPCRegistrationFunc func(server *grpc.Server)

func GRPCServer(o config.Observable, reg GRPCRegistrationFunc, opts ...grpc.ServerOption) *grpc.Server {
	if o.EnableObservability() {
		so := opentelemetry.ServerOption(opentelemetry.Options{
			MetricsOptions: opentelemetry.MetricsOptions{
				MeterProvider: otel.GetMeterProvider(),
				Metrics: opentelemetry.DefaultMetrics().Add(
					"grpc.lb.pick_first.connection_attempts_succeeded",
					"grpc.lb.pick_first.connection_attempts_failed",
				),
			},
			TraceOptions: oteltracing.TraceOptions{
				TracerProvider:    otel.GetTracerProvider(),
				TextMapPropagator: propagation.TraceContext{},
			},
		})
		opts = append(opts, so)
	}

	g := grpc.NewServer(opts...)

	reflection.Register(g)
	//TODO add health check apis
	reg(g)

	return g
}
