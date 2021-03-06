package database

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/docker/go-connections/nat"
	"github.com/google/wire"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"go.uber.org/zap"
	"test/internal/pkg/database"
	"time"
)

func createPostgres() (testcontainers.GenericContainerRequest, string, func(port nat.Port) string) {
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
	return req, port, dbURL
}

func NewSQlDb(ctx context.Context, o *database.Options, logger *zap.Logger) (*sql.DB, error) {
	var (
		req   testcontainers.GenericContainerRequest
		port  string
		dbURL func(port nat.Port) string
	)
	switch o.GetDialect() {
	case "postgres":
		req, port, dbURL = createPostgres()
	}
	container, err := testcontainers.GenericContainer(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to start container: %s", err)
	}
	mappedPort, err := container.MappedPort(ctx, nat.Port(port))
	if err != nil {
		return nil, fmt.Errorf("failed to get container external port: %s", err)
	}
	logger.Info(o.GetDialect()+" container ready and running at port: ", zap.Any("mappedPort", mappedPort))
	return sql.Open(o.GetDialect(), dbURL(mappedPort))
}

var ProviderSet = wire.NewSet(NewSQlDb, database.NewGormDb, database.NewOptions)
