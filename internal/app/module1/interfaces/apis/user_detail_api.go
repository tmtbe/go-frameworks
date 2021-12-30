package apis

import (
	cache "github.com/chenyahui/gin-cache"
	"github.com/gin-gonic/gin"
	"test/internal/app/module1/application"
	"test/internal/app/module1/interfaces/exceptions"
	"time"
)

type UserDetailAPI struct {
	API
	application *application.UserDetailApplication
}

func (dc *UserDetailAPI) GetRoute() {
	dc.ctx.Route.GET("/detail", cache.CacheByRequestURI(dc.ctx.CacheStore, 2*time.Second), wrapper(dc.GetUserDetail))
}

func NewUserDetailAPI(api *API, a *application.UserDetailApplication) *UserDetailAPI {
	return &UserDetailAPI{
		API:         *api,
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
