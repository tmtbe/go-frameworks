package apis

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"test/internal/app/module1/infrastructure/exceptions"
	exception2 "test/internal/app/module1/interfaces/exceptions"
	"test/internal/pkg/context"
	th "test/internal/pkg/transports/http"
)

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

func CreateInitControllersFn(
	pc ...th.Controller,
) th.InitControllers {
	return func(ctx *context.AppContext) {
		for _, c := range pc {
			c.GetRoute(ctx)
		}
	}
}
