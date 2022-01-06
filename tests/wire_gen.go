// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package tests

import (
	"github.com/google/wire"
	"test/internal/app"
	context2 "test/internal/app/context"
	"test/internal/app/module1/application"
	"test/internal/app/module1/domain/services"
	"test/internal/app/module1/infrastructure/repos"
	"test/internal/app/module1/interfaces/apis"
	"test/internal/pkg/cachestore"
	"test/internal/pkg/config"
	"test/internal/pkg/database"
	"test/internal/pkg/log"
	"test/internal/pkg/migrate"
	"test/tests/pkg"
	"test/tests/pkg/context"
	database2 "test/tests/pkg/database"
	"test/tests/pkg/redis"
	"test/tests/pkg/testcontainer"
	"test/tests/pkg/transports/http"
)

import (
	_ "github.com/lib/pq"
)

// Injectors from wire.go:

func CreateBackground(cf string) (*testcontainer.Background, func(), error) {
	viper, err := config.New(cf)
	if err != nil {
		return nil, nil, err
	}
	options, err := log.NewOptions(viper)
	if err != nil {
		return nil, nil, err
	}
	logger, err := log.New(options)
	if err != nil {
		return nil, nil, err
	}
	databaseOptions, err := database.NewOptions(viper, logger)
	if err != nil {
		return nil, nil, err
	}
	migrationOptions, err := migrate.NewOptions(viper)
	if err != nil {
		return nil, nil, err
	}
	contextContext := context.NewContext()
	db, err := database2.NewSQlDb(contextContext, databaseOptions, logger)
	if err != nil {
		return nil, nil, err
	}
	init, err := migrate.NewInit(viper, databaseOptions, migrationOptions, db, logger)
	if err != nil {
		return nil, nil, err
	}
	engine := http.NewTestGin(logger)
	gormDB, err := database.NewGormDb(db, logger)
	if err != nil {
		return nil, nil, err
	}
	client, err := redis.NewRedis(contextContext, logger)
	if err != nil {
		return nil, nil, err
	}
	redisStore := cachestore.NewRedisCache(client)
	testInfraContext := &context.TestInfraContext{
		MigrateInit: init,
		Config:      viper,
		Log:         logger,
		Route:       engine,
		GormDB:      gormDB,
		DB:          db,
		CacheStore:  redisStore,
		Context:     contextContext,
	}
	api := apis.NewAPI(logger, testInfraContext)
	postgresDetailRepository := repos.NewPostgresDetailsRepository(logger, gormDB)
	postgresUserRepository := repos.NewPostgresUserRepository(logger, gormDB)
	userDetailServiceImpl := services.NewUserDetailServiceImpl(logger, postgresDetailRepository, postgresUserRepository)
	userDetailApplication := application.NewUserDetailsApplication(logger, userDetailServiceImpl)
	userDetailAPI := apis.NewUserDetailAPI(api, userDetailApplication)
	appContext := &context2.AppContext{
		InfraContext:          testInfraContext,
		UserDetailAPI:         userDetailAPI,
		UserDetailApplication: userDetailApplication,
		UserRepository:        postgresUserRepository,
		DetailRepository:      postgresDetailRepository,
		UserDetailService:     userDetailServiceImpl,
	}
	background := &testcontainer.Background{
		AppContext:            appContext,
		TestContainersContext: contextContext,
	}
	return background, func() {
	}, nil
}

func CreateUserDetailAPI(cf string, s services.UserDetailService) (*apis.UserDetailAPI, error) {
	viper, err := config.New(cf)
	if err != nil {
		return nil, err
	}
	options, err := log.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	logger, err := log.New(options)
	if err != nil {
		return nil, err
	}
	engine := http.NewTestGin(logger)
	memoryStore := cachestore.NewMemoryCache()
	testMockAPIInfraContext := &context.TestMockAPIInfraContext{
		Config:     viper,
		Log:        logger,
		Route:      engine,
		CacheStore: memoryStore,
	}
	api := apis.NewAPI(logger, testMockAPIInfraContext)
	userDetailApplication := application.NewUserDetailsApplication(logger, s)
	userDetailAPI := apis.NewUserDetailAPI(api, userDetailApplication)
	return userDetailAPI, nil
}

// wire.go:

var ProviderSet = wire.NewSet(app.ProviderSet, pkg.ProviderSet)
