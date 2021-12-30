package telemetry

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.uber.org/zap"
	"google.golang.org/grpc/credentials"
	"log"
)

type Options struct {
	Enable      bool
	ServiceName string
	Endpoint    string
	Token       string
	Insecure    bool
}

func NewOptions(v *viper.Viper, logger *zap.Logger) (*Options, error) {
	var err error
	o := new(Options)
	if err = v.UnmarshalKey("telemetry", o); err != nil {
		return nil, errors.Wrap(err, "unmarshal telemetry option error")
	}
	logger.Info("load telemetry options success", zap.String("endpoint", o.Endpoint))
	return o, err
}

type Init struct {
}

func NewInit(ctx context.Context, o *Options, logger *zap.Logger, engine *gin.Engine) (*Init, func()) {
	if !o.Enable {
		return &Init{}, nil
	}
	engine.Use(otelgin.Middleware(o.ServiceName))
	secureOption := otlptracegrpc.WithTLSCredentials(credentials.NewClientTLSFromCert(nil, ""))
	if o.Insecure {
		secureOption = otlptracegrpc.WithInsecure()
	}
	headers := map[string]string{
		"signoz-access-token": o.Token,
	}
	exporter, err := otlptrace.New(
		context.Background(),
		otlptracegrpc.NewClient(
			secureOption,
			otlptracegrpc.WithEndpoint(o.Endpoint),
			otlptracegrpc.WithHeaders(headers),
		),
	)

	if err != nil {
		log.Fatal(err)
	}
	resources, err := resource.New(
		context.Background(),
		resource.WithAttributes(
			attribute.String("service.name", o.ServiceName),
			attribute.String("library.language", "go"),
		),
	)
	if err != nil {
		logger.Error("Could not set resources", zap.Error(err))
	}

	otel.SetTracerProvider(
		sdktrace.NewTracerProvider(
			sdktrace.WithSampler(sdktrace.AlwaysSample()),
			sdktrace.WithSpanProcessor(sdktrace.NewBatchSpanProcessor(exporter)),
			sdktrace.WithSyncer(exporter),
			sdktrace.WithResource(resources),
		),
	)
	return &Init{}, func() {
		exporter.Shutdown(ctx)
	}
}

var ProviderSet = wire.NewSet(NewInit, NewOptions)
