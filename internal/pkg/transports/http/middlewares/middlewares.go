package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"test/internal/pkg/telemetry"
)

type Middleware interface {
	GetMiddleware() gin.HandlerFunc
}

// NewMiddlewares 这里注册中间件
func NewMiddlewares(telemetryMiddleware *telemetry.Middleware) ([]Middleware, func()) {
	var middlewares = []Middleware{
		telemetryMiddleware,
	}
	return middlewares, nil
}

var ProviderSet = wire.NewSet(NewMiddlewares)
