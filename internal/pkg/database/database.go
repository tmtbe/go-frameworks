package database

import (
	"database/sql"
	"github.com/google/wire"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// Options is  configuration of database
type Options struct {
	URL   string
	Debug bool
}

func NewOptions(v *viper.Viper, logger *zap.Logger) (*Options, error) {
	var err error
	o := new(Options)
	if err = v.UnmarshalKey("db", o); err != nil {
		return nil, errors.Wrap(err, "unmarshal db option error")
	}

	logger.Info("load database options success", zap.String("url", o.URL))

	return o, err
}

// New Init 初始化数据库
func New(o *Options) (*sql.DB, error) {
	sqlDB, err := sql.Open("postgres", o.URL)
	if err != nil {
		return nil, errors.Wrap(err, "database open error")
	}
	return sqlDB, nil
}

var ProviderSet = wire.NewSet(New, NewOptions)
