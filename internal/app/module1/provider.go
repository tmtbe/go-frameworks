package module1

import (
	"github.com/google/wire"
	"test/internal/app/module1/application"
	"test/internal/app/module1/domain"
	"test/internal/app/module1/infrastructure"
	"test/internal/app/module1/interfaces"
)

var ProviderSet = wire.NewSet(
	application.ProviderSet,
	domain.ProviderSet,
	infrastructure.ProviderSet,
	interfaces.ProviderSet,
)
