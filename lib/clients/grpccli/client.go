package grpccli

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	oteltracing "google.golang.org/grpc/experimental/opentelemetry"
	"google.golang.org/grpc/stats/opentelemetry"
)

type Client struct {
	*grpc.ClientConn
}

func New(addr string) (*grpc.ClientConn, error) {
	return grpc.NewClient(addr, DialOpts()...)
}

func DialOpts() []grpc.DialOption {
	otelDialOpts := opentelemetry.DialOption(opentelemetry.Options{
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

	return []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		otelDialOpts,
	}
}
