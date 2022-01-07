package context

import (
	"github.com/google/wire"
	"test/gen/restapi"
	apis2 "test/internal/app/module1/adapter/apis"
	"test/internal/app/module1/adapter/repos"
	reposDef "test/internal/app/module1/domain/repos"
	"test/internal/app/module1/domain/services"
	"test/internal/app/module1/usercase"
	"test/internal/pkg/context"
)

type AppContext struct {
	Routes *restapi.Routes

	context.InfraContext

	*apis2.UserDetailAPI

	*usercase.UserDetailUsercase

	reposDef.UserRepository
	reposDef.DetailRepository

	services.UserDetailService
}

var ProviderSet = wire.NewSet(
	wire.Struct(new(AppContext), "*"),
	// API
	APIProviderSet,
	// Usercase
	UsercaseProviderSet,
	// Service
	ServiceProviderSet,
	// Repo
	RepoProviderSet,
)

var APIProviderSet = wire.NewSet(
	apis2.NewAPI,
	apis2.NewUserDetailAPI,
	// 自动生成的API Routes注入
	restapi.NewRoutes,
)

var UsercaseProviderSet = wire.NewSet(
	usercase.NewUserDetailsUsercase,
	usercase.NewHealthyUsercaseImpl,
	usercase.NewUserUsercaseImpl,
	usercase.NewPetUsercaseImpl,
	wire.Bind(new(restapi.UserUsercase), new(*usercase.UserUsercaseImpl)),
	wire.Bind(new(restapi.HealthyUsercase), new(*usercase.HealthyUsercaseImpl)),
	wire.Bind(new(restapi.PetUsercase), new(*usercase.PetUsercaseImpl)),
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
