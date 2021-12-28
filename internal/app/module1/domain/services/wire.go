//go:generate go run github.com/google/wire/cmd/wire
// +build wireinject

package services

import (
	"github.com/google/wire"
	"test/internal/app/module1/infrastructure"
	"test/internal/pkg"
)

var testProviderSet = wire.NewSet(
	ProviderSet,
	infrastructure.ProviderSet,
	pkg.ProviderSet,
)

func CreateUserDetailService(cf string) (UserDetailService, error) {
	panic(wire.Build(testProviderSet))
}
