package repo

import (
	"github.com/blackhorseya/godutch/internal/pkg/base/contextx"
	"github.com/google/wire"
)

// IRepo declare health repo service function
type IRepo interface {
	// Ping serve caller to get connect status
	Ping(ctx contextx.Contextx) (bool, error)
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(NewImpl)
