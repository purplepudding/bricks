package observability

import (
	"context"

	"go.opentelemetry.io/contrib/exporters/autoexport"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.40.0"
)

func EnableExport(serviceName, version string) (*sdktrace.TracerProvider, *sdkmetric.MeterProvider, error) {
	ctx := context.Background()

	r, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(serviceName),
			semconv.ServiceVersion(version),
		),
	)
	if err != nil {
		return nil, nil, err
	}

	tp, err := buildTracerProvider(ctx, r)
	if err != nil {
		return nil, nil, err
	}

	mp, err := buildMeterProvider(ctx, r)
	if err != nil {
		return nil, nil, err
	}

	return tp, mp, nil
}

func buildTracerProvider(ctx context.Context, r *resource.Resource) (*sdktrace.TracerProvider, error) {
	exp, err := autoexport.NewSpanExporter(ctx)
	if err != nil {
		return nil, err
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp),
		sdktrace.WithResource(r),
	)

	return tp, nil
}

func buildMeterProvider(ctx context.Context, r *resource.Resource) (*sdkmetric.MeterProvider, error) {
	exp, err := autoexport.NewMetricReader(ctx)
	if err != nil {
		return nil, err
	}

	mp := sdkmetric.NewMeterProvider(
		sdkmetric.WithReader(exp),
		sdkmetric.WithResource(r),
	)

	return mp, nil
}
