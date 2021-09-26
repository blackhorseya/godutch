package history

import (
	"github.com/blackhorseya/godutch/internal/app/godutch/biz/history/repo"
	"github.com/blackhorseya/godutch/internal/pkg/base/contextx"
	"github.com/blackhorseya/godutch/internal/pkg/entity/event"
	"github.com/bwmarrin/snowflake"
	"go.uber.org/zap"
)

type impl struct {
	logger *zap.Logger
	repo   repo.IRepo
	node   *snowflake.Node
}

// NewImpl serve caller to create an IBiz
func NewImpl(logger *zap.Logger, repo repo.IRepo, node *snowflake.Node) IBiz {
	return &impl{
		logger: logger.With(zap.String("type", "HistoryBiz")),
		repo:   repo,
		node:   node,
	}
}

func (i *impl) GetByID(ctx contextx.Contextx, id int64) (info *event.Record, err error) {
	// todo: 2021-09-26|20:20|Sean|impl me
	panic("implement me")
}

func (i *impl) List(ctx contextx.Contextx, actID int64, page, size int) (infos []*event.Record, err error) {
	// todo: 2021-09-26|20:20|Sean|impl me
	panic("implement me")
}

func (i *impl) NewRecord(ctx contextx.Contextx, created *event.Record) (info *event.Record, err error) {
	// todo: 2021-09-26|20:20|Sean|impl me
	panic("implement me")
}

func (i *impl) Delete(ctx contextx.Contextx, id, actID int64) error {
	// todo: 2021-09-26|20:20|Sean|impl me
	panic("implement me")
}
