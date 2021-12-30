package interfaces

import (
	"github.com/google/wire"
	"test/internal/app/module1/interfaces/apis"
	"test/internal/pkg/transports/http"
)

// NewAPIS 这里注册API
func NewAPIS(userDetailApi *apis.UserDetailAPI) []http.Controller {
	var controllers = []http.Controller{
		userDetailApi,
	}
	return controllers
}

var ProviderSet = wire.NewSet(
	NewAPIS, apis.NewAPI, apis.NewUserDetailAPI,
)
