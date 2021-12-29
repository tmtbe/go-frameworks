package apis

import (
	cache "github.com/chenyahui/gin-cache"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"test/internal/app/module1/application"
	"test/internal/app/module1/interfaces/exceptions"
	"test/internal/pkg/context"
	"time"
)

type UserDetailAPI struct {
	logger      *zap.Logger
	application *application.UserDetailApplication
}

func (dc *UserDetailAPI) GetRoute(ctx *context.AppContext) {
	ctx.Route.GET("/detail", cache.CacheByRequestURI(ctx.CacheStore, 2*time.Second), wrapper(dc.GetUserDetail))
}

func NewUserDetailAPI(logger *zap.Logger, a *application.UserDetailApplication) *UserDetailAPI {
	return &UserDetailAPI{
		logger:      logger,
		application: a,
	}
}

func (dc *UserDetailAPI) GetUserDetail(c *gin.Context) (interface{}, error) {
	param := struct {
		ID uint64 `form:"id" binding:"required"`
	}{}
	err := c.ShouldBindQuery(&param)
	if err != nil {
		return nil, exceptions.ParameterError(err.Error())
	}
	p, err := dc.application.GetUserDetail(param.ID)
	if err != nil {
		return nil, err
	}
	return p, nil
}
