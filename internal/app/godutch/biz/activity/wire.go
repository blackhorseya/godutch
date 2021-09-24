//go:build wireinject
// +build wireinject

package activity

import (
	"github.com/blackhorseya/godutch/internal/app/godutch/biz/activity/repo"
	"github.com/bwmarrin/snowflake"
	"github.com/google/wire"
	"go.uber.org/zap"
)

var testProviderSet = wire.NewSet(NewImpl)

// CreateIBiz serve caller to create an IBiz
func CreateIBiz(logger *zap.Logger, repo repo.IRepo, node *snowflake.Node) (IBiz, error) {
	panic(wire.Build(testProviderSet))
}
