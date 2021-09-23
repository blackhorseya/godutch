// +build wireinject

package user

import (
	"github.com/blackhorseya/godutch/internal/app/godutch/biz/user/repo"
	"github.com/blackhorseya/godutch/internal/pkg/infra/token"
	"github.com/bwmarrin/snowflake"
	"github.com/google/wire"
	"go.uber.org/zap"
)

var testProviderSet = wire.NewSet(NewImpl)

// CreateIBiz serve caller to create an IBiz
func CreateIBiz(logger *zap.Logger, repo repo.IRepo, node *snowflake.Node, token *token.Factory) (IBiz, error) {
	panic(wire.Build(testProviderSet))
}
