package monitoring

import (
	"context"
	"log"

	"github.com/ananrafs1/gomic/model"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/trace/jaeger"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

var tracer trace.Tracer

func Initialize() {
	exporter, err := jaeger.NewRawExporter(
		jaeger.WithCollectorEndpoint(model.Config.Monitoring["endPoint"]),
		jaeger.WithProcess(jaeger.Process{
			ServiceName: model.Config.Monitoring["serviceName"],
		}),
	)
	if err != nil {
		log.Fatalf("[Error] - %s ", err)
	}

	traceProvider := sdktrace.NewTracerProvider(
		sdktrace.WithSyncer(exporter))

	otel.SetTracerProvider(traceProvider)
	tracer = traceProvider.Tracer(model.Config.Monitoring["serviceName"])
}

func CreateSpan(ctx context.Context, process string) (trace.Span, func(), context.Context) {
	ctx, span := tracer.Start(ctx, process)
	return span, func() { span.End() }, ctx
}
