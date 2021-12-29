package app

import (
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
	"test/internal/pkg/transports/http"
	"test/internal/pkg/utils"
)

type Application struct {
	name       string
	logger     *zap.Logger
	httpServer *http.Server
	context    *Context
}

type Context struct {
	config  *viper.Viper
	log     *zap.Logger
	content map[string]interface{}
}

func (c *Context) Add(name string, component interface{}) {
	if _, ok := c.content[name]; ok {
		panic("has same key component")
	}
	c.content[name] = component
	c.log.Debug("context add component:", zap.Any(name, utils.Typeof(component)))
}

func (c *Context) Has(name string) (ok bool) {
	_, ok = c.content[name]
	return
}

func (c *Context) Get(name string) (interface{}, error) {
	if i, ok := c.content[name]; ok {
		return i, nil
	}
	return nil, errors.New("no component:" + name)
}

func (c *Context) MustGet(name string) interface{} {
	if i, ok := c.content[name]; ok {
		return i
	}
	panic("no component:" + name)
}

func NewAppContext(config *viper.Viper, log *zap.Logger) *Context {
	return &Context{
		config:  config,
		log:     log,
		content: map[string]interface{}{},
	}
}

type Option func(app *Application) error

func HttpServerOption(svr *http.Server) Option {
	return func(app *Application) error {
		svr.Application(app.name)
		app.httpServer = svr

		return nil
	}
}

func New(name string, context *Context, logger *zap.Logger, options ...Option) (*Application, error) {
	app := &Application{
		name:    name,
		logger:  logger.With(zap.String("type", "Application")),
		context: context,
	}

	for _, option := range options {
		if err := option(app); err != nil {
			return nil, err
		}
	}

	return app, nil
}

func (a *Application) Start() error {
	if a.httpServer != nil {
		if err := a.httpServer.Start(); err != nil {
			return errors.Wrap(err, "http server start error")
		}
	}
	return nil
}

func (a *Application) AwaitSignal() {
	c := make(chan os.Signal, 1)
	signal.Reset(syscall.SIGTERM, syscall.SIGINT)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)
	select {
	case s := <-c:
		a.logger.Info("receive a signal", zap.String("signal", s.String()))
		if a.httpServer != nil {
			if err := a.httpServer.Stop(); err != nil {
				a.logger.Warn("stop http server error", zap.Error(err))
			}
		}
		os.Exit(0)
	}
}

var ProviderSet = wire.NewSet(NewAppContext)
