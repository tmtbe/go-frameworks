package interfaces

import (
	"github.com/google/wire"
	"go.uber.org/zap"
	"test/internal/app/module1/application"
	"test/internal/app/module1/interfaces/apis"
	"test/internal/pkg/route"
)

// NewAPIS 这里注册API
func NewAPIS(logger *zap.Logger, a *application.UserDetailApplication) []route.Controller {
	var controllers []route.Controller
	controllers = append(controllers, apis.NewUserDetailAPI(logger, a))
	return controllers
}

var ProviderSet = wire.NewSet(
	NewAPIS,
	apis.CreateInitControllersFn,
)
