package migrate

import (
	"database/sql"
	"github.com/google/wire"
	"github.com/pkg/errors"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"test/internal/pkg/database"
)

type MigrationOptions struct {
	Dir string
}

func NewOptions(v *viper.Viper) (*MigrationOptions, error) {
	var err error
	o := new(MigrationOptions)
	if err = v.UnmarshalKey("db.migrations", o); err != nil {
		return nil, errors.Wrap(err, "unmarshal migration option error")
	}
	return o, err
}

type Init struct {
}

func NewInit(v *viper.Viper, o *database.Options, mo *MigrationOptions, sqlDb *sql.DB, logger *zap.Logger) (*Init, error) {
	if sqlDb == nil {
		logger.Warn("sql db is nil, migrate is skips")
		return nil, nil
	}
	m := &migrate.FileMigrationSource{
		Dir: v.GetString("resources_path") + mo.Dir,
	}
	n, err := migrate.Exec(sqlDb, o.GetDialect(), m, migrate.Up)
	if err != nil {
		return nil, errors.Wrap(err, "applying migrations failed")
	}
	logger.Info("migrations applied", zap.Int("count", n))
	return &Init{}, nil
}

var ProviderSet = wire.NewSet(NewOptions, NewInit)
