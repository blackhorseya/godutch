//go:build wireinject
// +build wireinject

package user

import (
	"github.com/blackhorseya/godutch/internal/app/godutch/biz/user"
	"github.com/google/wire"
	"go.uber.org/zap"
)

var testProviderSet = wire.NewSet(NewImpl)

// CreateIHandler serve caller to create an IHandler
func CreateIHandler(logger *zap.Logger, biz user.IBiz) (IHandler, error) {
	panic(wire.Build(testProviderSet))
}
