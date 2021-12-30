package apis

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	appContext "test/internal/app/context"
	"test/internal/app/module1/infrastructure/exceptions"
	exception2 "test/internal/app/module1/interfaces/exceptions"
)

type API struct {
	logger *zap.Logger
	ctx    *appContext.Context
}

func NewAPI(logger *zap.Logger, ctx *appContext.Context) *API {
	return &API{
		logger: logger,
		ctx:    ctx,
	}
}

type HandlerFunc func(c *gin.Context) (interface{}, error)

func wrapper(handler HandlerFunc) func(c *gin.Context) {
	return func(c *gin.Context) {
		var (
			err  error
			body interface{}
		)
		body, err = handler(c)
		if err != nil {
			var apiException *exceptions.AppException
			if h, ok := err.(*exceptions.AppException); ok {
				apiException = h
			} else if e, ok := err.(error); ok {
				if gin.Mode() == "debug" {
					// 错误
					apiException = exception2.UnknownError(e.Error())
				} else {
					// 未知错误
					apiException = exception2.UnknownError(e.Error())
				}
			} else {
				apiException = exception2.ServerError()
			}
			apiException.Request = c.Request.Method + " " + c.Request.URL.String()
			c.JSON(apiException.Code, apiException)
			return
		} else {
			c.JSON(http.StatusOK, body)
		}
	}
}
