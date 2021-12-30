package database

import (
	"database/sql"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"strings"
)

// Options is  configuration of database
type Options struct {
	URL    string
	Enable bool
}

func (o *Options) GetDialect() string {
	return strings.Split(o.URL, ":")[0]
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

// NewSqlDb Init 初始化数据库
func NewSqlDb(o *Options) (*sql.DB, error) {
	if !o.Enable {
		return nil, nil
	}
	sqlDB, err := sql.Open(o.GetDialect(), o.URL)
	if err != nil {
		return nil, errors.Wrap(err, "database open error")
	}
	return sqlDB, nil
}

func NewGormDb(sqlDb *sql.DB, logger *zap.Logger) (*gorm.DB, error) {
	if sqlDb == nil {
		logger.Warn("sql db is nil, gorm db will not create")
		return nil, nil
	}
	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDb,
	}), &gorm.Config{})
	if err != nil {
		return nil, errors.Wrap(err, "database open error")
	}
	return db, nil
}

var ProviderSet = wire.NewSet(NewSqlDb, NewGormDb, NewOptions)
