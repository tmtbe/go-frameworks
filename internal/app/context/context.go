package context

import (
	"github.com/google/wire"
	"test/internal/app/module1/application"
	services2 "test/internal/app/module1/domain/expose"
	"test/internal/app/module1/domain/services"
	repos2 "test/internal/app/module1/infrastructure/expose"
	"test/internal/app/module1/infrastructure/repos"
	"test/internal/app/module1/interfaces/apis"
	"test/internal/pkg/context"
)

type AppContext struct {
	context.InfraContext

	*apis.UserDetailAPI

	*application.UserDetailApplication

	repos2.UserRepository
	repos2.DetailRepository

	services2.UserDetailService
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
	wire.Bind(new(services2.UserDetailService), new(*services.UserDetailServiceImpl)),
)

var RepoProviderSet = wire.NewSet(
	repos.NewPostgresDetailsRepository,
	repos.NewPostgresUserRepository,
	wire.Bind(new(repos2.UserRepository), new(*repos.PostgresUserRepository)),
	wire.Bind(new(repos2.DetailRepository), new(*repos.PostgresDetailRepository)),
)
