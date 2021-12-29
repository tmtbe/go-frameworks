package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Options struct {
	Network string
	URL     string
	DB      int
	Enable  bool
}

func NewOptions(v *viper.Viper, logger *zap.Logger) (*Options, error) {
	var err error
	o := new(Options)
	if err = v.UnmarshalKey("redis", o); err != nil {
		return nil, errors.Wrap(err, "unmarshal redis option error")
	}
	logger.Info("load redis options success", zap.String("url", o.URL))
	return o, err
}

// New Init 初始化Redis
func New(ctx context.Context, o *Options) (*redis.Client, error) {
	if !o.Enable {
		return nil, nil
	}
	client := redis.NewClient(&redis.Options{
		Network: o.Network,
		Addr:    o.URL,
		DB:      o.DB,
	})
	ping := client.Ping(ctx)
	if ping.Err() != nil {
		return nil, errors.Wrap(ping.Err(), "redis connect error")
	}
	return client, nil
}

var ProviderSet = wire.NewSet(New, NewOptions)
