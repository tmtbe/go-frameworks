package context

import (
	"context"
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"test/internal/pkg/cachestore"
	"test/internal/pkg/migrate"
	"test/internal/pkg/telemetry"
)

func NewContext() context.Context {
	return context.Background()
}

// AppContext Init需要放在Context里面承载，不然会被忽略
type AppContext struct {
	MigrateInit   *migrate.Init
	TelemetryInit *telemetry.Init
	Config        *viper.Viper
	Log           *zap.Logger
	Route         *gin.Engine
	GormDB        *gorm.DB
	DB            *sql.DB
	CacheStore    cachestore.Store
	Context       context.Context
}

var ProviderSet = wire.NewSet(
	NewContext,
	wire.Struct(new(AppContext), "*"),
)
