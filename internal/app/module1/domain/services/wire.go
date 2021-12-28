//go:generate go run github.com/google/wire/cmd/wire
// +build wireinject

package services

import (
	"context"
	"github.com/google/wire"
	"gorm.io/gorm"
	"test/internal/app/module1/infrastructure"
	"test/internal/pkg"
)

var testProviderSet = wire.NewSet(
	ProviderSet,
	infrastructure.ProviderSet,
	pkg.TestProviderSet,
	NewTestContext,
)

type TestContext struct {
	userDetailService     UserDetailService
	db                    *gorm.DB
	testContainersContext context.Context
}

func NewTestContext(userDetailService UserDetailService, db *gorm.DB) *TestContext {
	return &TestContext{
		userDetailService:     userDetailService,
		db:                    db,
		testContainersContext: context.Background(),
	}
}

func CreateTestContext(cf string) (*TestContext, error) {
	panic(wire.Build(testProviderSet))
}
