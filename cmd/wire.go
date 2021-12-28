//go:generate go run github.com/google/wire/cmd/wire
// +build wireinject

package main

import (
	"github.com/google/wire"
	"test/internal/app"
	"test/internal/pkg"
	pkg_app "test/internal/pkg/app"
)

var providerSet = wire.NewSet(
	pkg.ProviderSet,
	app.ProviderSet,
)

func CreateApp(cf string) (*pkg_app.Application, error) {
	panic(wire.Build(providerSet))
}
