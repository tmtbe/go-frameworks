package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"go.uber.org/zap"
)

func NewRedis(ctx context.Context, logger *zap.Logger) (*redis.Client, error) {
	req := testcontainers.ContainerRequest{
		Image:        "redis:6",
		ExposedPorts: []string{"6379/tcp"},
		WaitingFor:   wait.ForLog("* Ready to accept connections"),
	}
	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return nil, err
	}
	mappedPort, err := container.MappedPort(ctx, "6379")
	if err != nil {
		return nil, err
	}
	hostIP, err := container.Host(ctx)
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s:%s", hostIP, mappedPort.Port())
	client := redis.NewClient(&redis.Options{
		Network: "tcp",
		Addr:    url,
	})
	ping := client.Ping(ctx)
	if ping.Err() != nil {
		return nil, errors.Wrap(ping.Err(), "redis connect error")
	}
	logger.Info("redis container ready and running at port: ", zap.Any("mappedPort", mappedPort))
	return client, nil
}

var ProviderSet = wire.NewSet(NewRedis)
