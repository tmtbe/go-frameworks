package pkg

import (
	"github.com/google/wire"
	"test/internal/pkg/app"
	"test/internal/pkg/config"
	"test/internal/pkg/database"
	"test/internal/pkg/log"
	"test/internal/pkg/migrate"
	"test/internal/pkg/tests"
	"test/internal/pkg/transports/http"
)

var ProviderSet = wire.NewSet(
	app.ProviderSet,
	log.ProviderSet,
	config.ProviderSet,
	database.ProviderSet,
	migrate.ProviderSet,
	http.ProviderSet,
)
var TestProviderSet = wire.NewSet(
	app.ProviderSet,
	log.ProviderSet,
	config.ProviderSet,
	tests.ProviderSet,
	migrate.ProviderSet,
	http.ProviderSet,
)
