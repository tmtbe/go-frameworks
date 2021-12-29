package testcontainer

import (
	"context"
	"github.com/google/wire"
	app_context "test/internal/app/context"
)

func NewTSContext() context.Context {
	background := context.Background()
	return background
}

type Background struct {
	app_context.Context
	TestContainersContext context.Context
}

var ProviderSet = wire.NewSet(NewTSContext, wire.Struct(new(Background), "*"))
