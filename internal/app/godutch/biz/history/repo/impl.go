package repo

import (
	"time"

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
		rw:     rw,
	}
}

func (i *impl) GetByID(ctx contextx.Contextx, id int64) (info *event.Record, err error) {
	timeout, cancel := contextx.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var ret event.Record
	stmt := `
select h.id         as id,
       h.remark     as remark,
       user.id      as "payer.id",
       user.email   as "payer.email",
       user.name    as "payer.name",
       detail.value as total,
       h.created_at as created_at
from spend_history h
         join users user on user.id = h.payer_id
         join spend_details detail on h.payer_id = detail.user_id
where h.id = ? 
order by detail.value desc limit 1`
	err = i.rw.GetContext(timeout, &ret, stmt, id)
	if err != nil {
		return nil, err
	}

	return &ret, nil
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
