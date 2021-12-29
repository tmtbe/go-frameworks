package infrastructure

import (
	"github.com/google/wire"
	"test/internal/app/module1/infrastructure/repos"
)

// ProviderSet 注册所有的repo
var ProviderSet = wire.NewSet(
	repos.NewPostgresDetailsRepository,
	repos.NewPostgresUserRepository,
	wire.Bind(new(repos.UserRepository), new(*repos.PostgresUserRepository)),
	wire.Bind(new(repos.DetailRepository), new(*repos.PostgresDetailRepository)),
)
