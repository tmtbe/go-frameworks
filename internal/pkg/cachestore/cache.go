package cachestore

import (
	"github.com/chenyahui/gin-cache/persist"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"time"
)

type Store interface {
	Set(key string, value interface{}, expire time.Duration) error
	Delete(key string) error
	Get(key string, value interface{}) error
}

func NewRedisCache(client *redis.Client) *persist.RedisStore {
	return persist.NewRedisStore(client)
}

func NewMemoryCache() *persist.MemoryStore {
	return persist.NewMemoryStore(1)
}

var ProviderSetRedis = wire.NewSet(NewRedisCache, wire.Bind(new(Store), new(*persist.RedisStore)))
var ProviderSetMemory = wire.NewSet(NewMemoryCache, wire.Bind(new(Store), new(*persist.MemoryStore)))
