package context

import (
	"github.com/google/wire"
	"test/internal/app/module1/application"
	reposDef "test/internal/app/module1/domain/repos"
	"test/internal/app/module1/domain/services"
	"test/internal/app/module1/infrastructure/repos"
	"test/internal/app/module1/interfaces/apis"
	"test/internal/gen/restapi"
	"test/internal/pkg/context"
)

type AppContext struct {
	Routes *restapi.Routes

	context.InfraContext

	*apis.UserDetailAPI

	*application.UserDetailApplication

	reposDef.UserRepository
	reposDef.DetailRepository

	services.UserDetailService
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
	// 自动生成的API Routes注入
	restapi.NewRoutes,
)

var ApplicationProviderSet = wire.NewSet(
	application.NewUserDetailsApplication,
	application.NewHealthyApplicationImpl,
	application.NewUserApplicationImpl,
	wire.Bind(new(restapi.UserApplication), new(*application.UserApplicationImpl)),
	wire.Bind(new(restapi.HealthyApplication), new(*application.HealthyApplicationImpl)),
)

var ServiceProviderSet = wire.NewSet(
	services.NewUserDetailServiceImpl,
	wire.Bind(new(services.UserDetailService), new(*services.UserDetailServiceImpl)),
)

var RepoProviderSet = wire.NewSet(
	repos.NewPostgresDetailsRepository,
	repos.NewPostgresUserRepository,
	wire.Bind(new(reposDef.UserRepository), new(*repos.PostgresUserRepository)),
	wire.Bind(new(reposDef.DetailRepository), new(*repos.PostgresDetailRepository)),
)
