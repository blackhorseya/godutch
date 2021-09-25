//go:build wireinject
// +build wireinject

package repo

import (
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

var testProviderSet = wire.NewSet(NewImpl)

// CreateRepo serve caller to create an IRepo
func CreateRepo(logger *zap.Logger, rw *sqlx.DB) (IRepo, error) {
	panic(wire.Build(testProviderSet))
}
