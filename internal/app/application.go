package app

import (
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"test/internal/app/module1"
	"test/internal/pkg/app"
	"test/internal/pkg/transports/http"
)

type Options struct {
	Name string
}

func NewOptions(v *viper.Viper, logger *zap.Logger) (*Options, error) {
	var err error
	o := new(Options)
	if err = v.UnmarshalKey("app", o); err != nil {
		return nil, errors.Wrap(err, "unmarshal app option error")
	}
	logger.Info("load application options success")
	return o, err
}

func NewApp(o *Options, context *app.Context, logger *zap.Logger, hs *http.Server) (*app.Application, error) {
	a, err := app.New(o.Name, context, logger, app.HttpServerOption(hs))
	if err != nil {
		return nil, errors.Wrap(err, "new app error")
	}
	return a, nil
}

var ProviderSet = wire.NewSet(
	NewApp,
	NewOptions,
	module1.ProviderSet,
)
