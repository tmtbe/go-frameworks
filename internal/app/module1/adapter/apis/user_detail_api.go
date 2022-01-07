package apis

import (
	cache "github.com/chenyahui/gin-cache"
	"github.com/gin-gonic/gin"
	"test/internal/app/module1/usercase"
	"test/internal/pkg/app"
	"time"
)

type UserDetailAPI struct {
	API
	application *usercase.UserDetailUsercase
}

func NewUserDetailAPI(api *API, a *usercase.UserDetailUsercase) *UserDetailAPI {
	v := &UserDetailAPI{
		API:         *api,
		application: a,
	}
	v.Init()
	return v
}

func (dc *UserDetailAPI) Init() {
	dc.ctx.GetRoute().GET("/detail", cache.CacheByRequestURI(dc.ctx.GetCacheStore(), 2*time.Second), wrapper(dc.GetUserDetail))
}

func (dc *UserDetailAPI) GetUserDetail(c *gin.Context) (interface{}, error) {
	param := struct {
		ID uint64 `form:"id" binding:"required"`
	}{}
	err := c.ShouldBindQuery(&param)
	if err != nil {
		return nil, app.ParameterError(err.Error())
	}
	p, err := dc.application.GetUserDetail(param.ID)
	if err != nil {
		return nil, err
	}
	return p, nil
}
