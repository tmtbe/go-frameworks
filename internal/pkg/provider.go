package pkg

import (
	"github.com/google/wire"
	"test/internal/pkg/context"
)

var ProviderSet = wire.NewSet(
	context.ProviderSet,
)
