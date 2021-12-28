package tests

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/docker/go-connections/nat"
	"github.com/google/wire"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"go.uber.org/zap"
	"time"
)

func NewTSContext() context.Context {
	return context.Background()
}
func NewDb(context context.Context, logger *zap.Logger) (*sql.DB, error) {
	dbname := "test"
	var env = map[string]string{
		"POSTGRES_PASSWORD": "root",
		"POSTGRES_USER":     "root",
		"POSTGRES_DB":       dbname,
	}
	var port = "5432/tcp"
	dbURL := func(port nat.Port) string {
		return fmt.Sprintf("postgres://root:root@localhost:%s/%s?sslmode=disable", port.Port(), dbname)
	}

	req := testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image:        "postgres:latest",
			ExposedPorts: []string{port},
			Cmd:          []string{"postgres", "-c", "fsync=off"},
			Env:          env,
			WaitingFor:   wait.ForSQL(nat.Port(port), "postgres", dbURL).Timeout(time.Second * 5),
		},
		Started: true,
	}
	container, err := testcontainers.GenericContainer(context, req)
	if err != nil {
		return nil, fmt.Errorf("failed to start container: %s", err)
	}

	mappedPort, err := container.MappedPort(context, nat.Port(port))
	if err != nil {
		return nil, fmt.Errorf("failed to get container external port: %s", err)
	}
	logger.Info("postgres container ready and running at port: ", zap.Any("mappedPort", mappedPort))
	sqlDB, err := sql.Open("postgres", dbURL(mappedPort))
	if err != nil {
		return nil, errors.Wrap(err, "database open error")
	}
	return sqlDB, nil
}

var ProviderSet = wire.NewSet(NewTSContext, NewDb)
