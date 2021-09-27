//go:build wireinject
// +build wireinject

package history

import (
	"github.com/blackhorseya/godutch/internal/app/godutch/biz/history"
	"github.com/google/wire"
	"go.uber.org/zap"
)

var testProviderSet = wire.NewSet(NewImpl)

// CreateIHandler serve caller to create an IHandler
func CreateIHandler(logger *zap.Logger, biz history.IBiz) (IHandler, error) {
	panic(wire.Build(testProviderSet))
}
