package repo

import (
	"github.com/blackhorseya/godutch/internal/pkg/base/contextx"
	"github.com/blackhorseya/godutch/internal/pkg/entity/event"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type impl struct {
	logger *zap.Logger
	rw     *sqlx.DB
}

// NewImpl serve caller to create an IRepo
func NewImpl(logger *zap.Logger, rw *sqlx.DB) IRepo {
	return &impl{
		logger: logger.With(zap.String("type", "HistoryRepo")),
		rw: rw,
	}
}

func (i *impl) GetByID(ctx contextx.Contextx, id int64) (info *event.Record, err error) {
	// todo: 2021-09-26|03:00|Sean|impl me
	panic("implement me")
}

func (i *impl) List(ctx contextx.Contextx, actID int64, limit, offset int) (infos []*event.Record, err error) {
	// todo: 2021-09-26|03:00|Sean|impl me
	panic("implement me")
}

func (i *impl) Create(ctx contextx.Contextx, created *event.Record) (info *event.Record, err error) {
	// todo: 2021-09-26|03:00|Sean|impl me
	panic("implement me")
}

func (i *impl) Delete(ctx contextx.Contextx, id int64) error {
	// todo: 2021-09-26|03:00|Sean|impl me
	panic("implement me")
}
