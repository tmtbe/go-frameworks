//go:generate go run github.com/google/wire/cmd/wire
// +build wireinject

package tests

import (
	"github.com/google/wire"
	"test/internal/app"
	"test/internal/app/context"
	"test/internal/app/module1/domain/expose"
	"test/internal/app/module1/interfaces/apis"
	"test/tests/pkg"
	"test/tests/pkg/testcontainer"
)

var ProviderSet = wire.NewSet(
	app.ProviderSet,
	pkg.ProviderSet,
)

func CreateBackground(cf string) (*testcontainer.Background, func(), error) {
	panic(wire.Build(ProviderSet))
}

func CreateUserDetailAPI(cf string, s expose.UserDetailService) (*apis.UserDetailAPI, error) {
	panic(wire.Build(pkg.APIMockProviderSet, context.APIProviderSet, context.ApplicationProviderSet))
}
