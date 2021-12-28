package pkg

import (
	"github.com/google/wire"
	"test/internal/pkg/config"
	"test/internal/pkg/database"
	"test/internal/pkg/log"
	"test/internal/pkg/transports/http"
)

var ProviderSet = wire.NewSet(
	log.ProviderSet,
	config.ProviderSet,
	database.ProviderSet,
	http.ProviderSet,
)
