package app

import (
	"context"
	"log"

	"github.com/alextanhongpin/pkg/grace"
	"go.opentelemetry.io/otel/label"

	"go.opentelemetry.io/otel/exporters/trace/jaeger"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

const (
	collectorEndpoint = "http://localhost:14268/api/traces"
	serviceName       = "graphql-server"
)

func NewTracer(sg *grace.ShutdownGroup) {
	_, flush, err := jaeger.NewExportPipeline(
		jaeger.WithCollectorEndpoint(collectorEndpoint),
		jaeger.WithProcess(jaeger.Process{
			ServiceName: serviceName,
			Tags: []label.KeyValue{
				label.String("exporter", "jaeger"),
			},
		}),
		jaeger.WithSDK(&sdktrace.Config{
			DefaultSampler: sdktrace.AlwaysSample(),
		}),
	)
	if err != nil {
		log.Fatal(err)
	}
	sg.Add(func(ctx context.Context) {
		flush()
	})
}
