package context

import (
	"github.com/google/wire"
	"test/internal/app/module1/application"
	servicesDef "test/internal/app/module1/application/services"
	reposDef "test/internal/app/module1/domain/repos"
	"test/internal/app/module1/domain/services"
	"test/internal/app/module1/infrastructure/repos"
	"test/internal/app/module1/interfaces/apis"
	"test/internal/pkg/context"
)

type AppContext struct {
	context.InfraContext

	*apis.UserDetailAPI

	*application.UserDetailApplication

	reposDef.UserRepository
	reposDef.DetailRepository

	servicesDef.UserDetailService
}

var ProviderSet = wire.NewSet(
	wire.Struct(new(AppContext), "*"),
	// API
	APIProviderSet,
	// Application
	ApplicationProviderSet,
	// Service
	ServiceProviderSet,
	// Repo
	RepoProviderSet,
)

var APIProviderSet = wire.NewSet(
	apis.NewAPI,
	apis.NewUserDetailAPI,
)

var ApplicationProviderSet = wire.NewSet(
	application.NewUserDetailsApplication,
)

var ServiceProviderSet = wire.NewSet(
	services.NewUserDetailServiceImpl,
	wire.Bind(new(servicesDef.UserDetailService), new(*services.UserDetailServiceImpl)),
)

var RepoProviderSet = wire.NewSet(
	repos.NewPostgresDetailsRepository,
	repos.NewPostgresUserRepository,
	wire.Bind(new(reposDef.UserRepository), new(*repos.PostgresUserRepository)),
	wire.Bind(new(reposDef.DetailRepository), new(*repos.PostgresDetailRepository)),
)
