package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

type Middleware interface {
	GetMiddleware() gin.HandlerFunc
}

// NewMiddlewares 这里注册中间件
func NewMiddlewares() ([]Middleware, func()) {
	var middlewares = []Middleware{}
	return middlewares, nil
}

var ProviderSet = wire.NewSet(NewMiddlewares)
