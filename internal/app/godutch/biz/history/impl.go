package history

import (
	"time"

	"github.com/blackhorseya/godutch/internal/app/godutch/biz/history/repo"
	"github.com/blackhorseya/godutch/internal/pkg/base/contextx"
	"github.com/blackhorseya/godutch/internal/pkg/entity/er"
	"github.com/blackhorseya/godutch/internal/pkg/entity/event"
	"github.com/blackhorseya/godutch/internal/pkg/entity/user"
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
	if len(ret) == 0 {
		i.logger.Error(er.ErrRecordNotExists.Error(), zap.Int64("act_id", actID), zap.Int("page", page), zap.Int("size", size))
		return nil, er.ErrRecordNotExists
	}

	return ret, nil
}

func (i *impl) NewRecord(ctx contextx.Contextx, actID, payerID int64, remark string, members []*user.Member, total int) (info *event.Record, err error) {
	if len(remark) == 0 {
		i.logger.Error(er.ErrEmptyRemark.Error())
		return nil, er.ErrEmptyRemark
	}

	if payerID == 0 {
		i.logger.Error(er.ErrMissingPayerID.Error())
		return nil, er.ErrMissingPayerID
	}

	if total == 0 {
		i.logger.Error(er.ErrMissingTotal.Error())
		return nil, er.ErrMissingTotal
	}

	created := &event.Record{
		ID:        i.node.Generate().Int64() / 1000 * 1000,
		Activity:  &event.Activity{ID: actID},
		Remark:    remark,
		Payer:     &user.Member{Id: payerID},
		Members:   members,
		Total:     total,
		CreatedAt: time.Now().Unix(),
	}
	ret, err := i.repo.Create(ctx, created)
	if err != nil {
		i.logger.Error(er.ErrNewRecord.Error(), zap.Error(err), zap.Any("created", created))
		return nil, er.ErrNewRecord
	}

	return ret, nil
}

func (i *impl) Delete(ctx contextx.Contextx, id, actID int64) error {
	err := i.repo.Delete(ctx, id)
	if err != nil {
		i.logger.Error(er.ErrDeleteRecord.Error(), zap.Error(err), zap.Int64("id", id))
		return er.ErrDeleteRecord
	}

	return nil
}
