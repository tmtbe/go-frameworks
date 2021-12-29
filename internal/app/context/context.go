package context

import (
	"github.com/google/wire"
	"test/internal/app/module1/domain/services"
	"test/internal/app/module1/infrastructure/repos"
	"test/internal/pkg/app"
)

type Context struct {
	*app.Context
	repos.UserRepository
	repos.DetailRepository
	services.UserDetailService
}

var ProviderSet = wire.NewSet(
	wire.Struct(new(Context), "*"),
)
