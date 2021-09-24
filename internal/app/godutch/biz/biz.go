package biz

import (
	"github.com/blackhorseya/godutch/internal/app/godutch/biz/activity"
	"github.com/blackhorseya/godutch/internal/app/godutch/biz/health"
	"github.com/blackhorseya/godutch/internal/app/godutch/biz/user"
	"github.com/google/wire"
)

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(
	health.ProviderSet,
	user.ProviderSet,
	activity.ProviderSet,
)
