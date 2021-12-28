package infrastructure

import (
	"github.com/google/wire"
	repos2 "test/internal/app/module1/infrastructure/repos"
)

// ProviderSet 注册所有的repo
var ProviderSet = wire.NewSet(
	repos2.NewPostgresDetailsRepository,
	repos2.NewPostgresUserRepository,
)
