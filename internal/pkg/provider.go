package pkg

import (
	"github.com/google/wire"
	"test/internal/pkg/cachestore"
	"test/internal/pkg/config"
	"test/internal/pkg/context"
	"test/internal/pkg/database"
	"test/internal/pkg/log"
	"test/internal/pkg/migrate"
	"test/internal/pkg/redis"
	"test/internal/pkg/transports/http"
)

var ProviderSet = wire.NewSet(
	context.ProviderSet,
	log.ProviderSet,
	config.ProviderSet,
	database.ProviderSet,
	redis.ProviderSet,
	migrate.ProviderSet,
	http.ProviderSet,
	cachestore.ProviderSetRedis,
)
