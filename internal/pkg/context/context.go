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
)

func NewContext() context.Context {
	return context.Background()
}

type AppContext struct {
	Config     *viper.Viper
	Log        *zap.Logger
	Route      *gin.Engine
	GormDB     *gorm.DB
	DB         *sql.DB
	CacheStore cachestore.Store
	Context    context.Context
}

var ProviderSet = wire.NewSet(NewContext, wire.Struct(new(AppContext), "*"))
