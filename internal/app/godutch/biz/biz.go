package biz

import (
	"github.com/blackhorseya/godutch/internal/app/godutch/biz/health"
	"github.com/google/wire"
)

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(
	health.ProviderSet,
)
