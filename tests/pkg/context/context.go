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
	"test/internal/pkg/config"
	context2 "test/internal/pkg/context"
	"test/internal/pkg/log"
	"test/internal/pkg/migrate"
	"test/internal/pkg/transports/http"
	"test/tests/pkg/database"
	"test/tests/pkg/redis"
)

func NewContext() context.Context {
	return context.Background()
}

// TestInfraContext Init需要放在Context里面承载，不然会被忽略
type TestInfraContext struct {
	MigrateInit *migrate.Init
	Config      *viper.Viper
	Log         *zap.Logger
	Route       *gin.Engine
	GormDB      *gorm.DB
	DB          *sql.DB
	CacheStore  cachestore.Store
	Context     context.Context
}

func (a *TestInfraContext) GetConfig() *viper.Viper {
	return a.Config
}
func (a *TestInfraContext) GetLog() *zap.Logger {
	return a.Log
}
func (a *TestInfraContext) GetRoute() *gin.Engine {
	return a.Route
}
func (a *TestInfraContext) GetGormDB() *gorm.DB {
	return a.GormDB
}
func (a *TestInfraContext) GetDB() *sql.DB {
	return a.DB
}
func (a *TestInfraContext) GetCacheStore() cachestore.Store {
	return a.CacheStore
}
func (a *TestInfraContext) GetContext() context.Context {
	return a.Context
}

var ProviderSet = wire.NewSet(
	NewContext,
	wire.Struct(new(TestInfraContext), "*"),
	wire.Bind(new(context2.InfraContext), new(*TestInfraContext)),
	log.ProviderSet,
	config.ProviderSet,
	database.ProviderSet,
	redis.ProviderSet,
	migrate.ProviderSet,
	http.ProviderSet,
	cachestore.ProviderSetRedis,
)
