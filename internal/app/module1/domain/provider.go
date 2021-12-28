package domain

import (
	"github.com/google/wire"
	"test/internal/app/module1/domain/services"
)

var ProviderSet = wire.NewSet(
	services.ProviderSet,
)
