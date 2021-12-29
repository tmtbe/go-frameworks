package migrate

import (
	"database/sql"
	"github.com/google/wire"
	"github.com/pkg/errors"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"test/internal/pkg/app"
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

func Migrate(c *app.Context, v *viper.Viper, o *MigrationOptions, sqlDb *sql.DB, logger *zap.Logger) (*gorm.DB, error) {
	m := &migrate.FileMigrationSource{
		Dir: v.GetString("resources_path") + o.Dir,
	}
	n, err := migrate.Exec(sqlDb, "postgres", m, migrate.Up)
	if err != nil {
		return nil, errors.Wrap(err, "applying migrations failed")
	}
	logger.Info("migrations applied", zap.Int("count", n))
	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDb,
	}), &gorm.Config{})
	if err != nil {
		return nil, errors.Wrap(err, "database open error")
	}
	c.Add("gorm_db", db)
	return db, nil
}

var ProviderSet = wire.NewSet(Migrate, NewOptions)
