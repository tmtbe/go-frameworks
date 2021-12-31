package http

import (
	"github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"go.uber.org/zap"
	"time"
)

func NewTestGin(logger *zap.Logger) *gin.Engine {
	// 配置gin
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(gin.Recovery()) // panic之后自动恢复
	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(logger, true))
	return r
}

var ProviderSet = wire.NewSet(NewTestGin)
