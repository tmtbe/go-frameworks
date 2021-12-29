package context

import (
	"github.com/google/wire"
	"test/internal/app/module1/domain/services"
	"test/internal/app/module1/infrastructure/repos"
	"test/internal/pkg/context"
)

type Context struct {
	*context.AppContext
	repos.UserRepository
	repos.DetailRepository
	services.UserDetailService
}

var ProviderSet = wire.NewSet(
	wire.Struct(new(Context), "*"),
)
