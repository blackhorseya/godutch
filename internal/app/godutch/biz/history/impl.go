package history

import (
	"github.com/blackhorseya/godutch/internal/app/godutch/biz/history/repo"
	"github.com/blackhorseya/godutch/internal/pkg/base/contextx"
	"github.com/blackhorseya/godutch/internal/pkg/entity/er"
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
	ret, err := i.repo.GetByID(ctx, id)
	if err != nil {
		i.logger.Error(er.ErrGetRecordByID.Error(), zap.Error(err), zap.Int64("id", id))
		return nil, er.ErrGetRecordByID
	}
	if ret == nil {
		i.logger.Error(er.ErrRecordNotExists.Error(), zap.Int64("id", id))
		return nil, er.ErrRecordNotExists
	}

	return ret, nil
}

func (i *impl) List(ctx contextx.Contextx, actID int64, page, size int) (infos []*event.Record, err error) {
	if page <= 0 {
		i.logger.Error(er.ErrInvalidPage.Error(), zap.Int("page", page))
		return nil, er.ErrInvalidPage
	}

	if size <= 0 {
		i.logger.Error(er.ErrInvalidSize.Error(), zap.Int("size", size))
		return nil, er.ErrInvalidSize
	}

	ret, err := i.repo.List(ctx, actID, size, (page-1)*size)
	if err != nil {
		i.logger.Error(er.ErrListRecords.Error(), zap.Error(err), zap.Int64("act_id", actID), zap.Int("page", page), zap.Int("size", size))
		return nil, er.ErrListRecords
	}
	if len(ret)  == 0 {
		i.logger.Error(er.ErrRecordNotExists.Error(), zap.Int64("act_id", actID), zap.Int("page", page), zap.Int("size", size))
		return nil, er.ErrRecordNotExists
	}

	return ret, nil
}

func (i *impl) NewRecord(ctx contextx.Contextx, created *event.Record) (info *event.Record, err error) {
	// todo: 2021-09-26|20:20|Sean|impl me
	panic("implement me")
}

func (i *impl) Delete(ctx contextx.Contextx, id, actID int64) error {
	// todo: 2021-09-26|20:20|Sean|impl me
	panic("implement me")
}
