package database

import (
	"github.com/google/wire"
	"github.com/pkg/errors"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type migrationOptions struct {
	Dir string
}

// Options is  configuration of database
type Options struct {
	URL        string
	Migrations migrationOptions
	Debug      bool
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
func New(o *Options, logger *zap.Logger) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  o.URL,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})
	if err != nil {
		return nil, errors.Wrap(err, "database open error")
	}
	sqlDb, err := db.DB()
	if err != nil {
		return nil, errors.Wrap(err, "database open error")
	}
	m := &migrate.FileMigrationSource{
		Dir: o.Migrations.Dir,
	}
	n, err := migrate.Exec(sqlDb, "postgres", m, migrate.Up)
	if err != nil {
		return nil, errors.Wrap(err, "applying migrations failed")
	}
	logger.Info("migrations applied", zap.Int("count", n))
	return db, nil
}

var ProviderSet = wire.NewSet(New, NewOptions)
