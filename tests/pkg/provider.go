package pkg

import (
	"github.com/google/wire"
	"test/internal/pkg/app"
	"test/internal/pkg/config"
	"test/internal/pkg/log"
	"test/internal/pkg/migrate"
	"test/internal/pkg/transports/http"
	"test/tests/pkg/database"
	"test/tests/pkg/testcontainer"
)

var ProviderSet = wire.NewSet(
	app.ProviderSet,
	log.ProviderSet,
	config.ProviderSet,
	database.ProviderSet,
	migrate.ProviderSet,
	http.ProviderSet,
	testcontainer.ProviderSet,
)
