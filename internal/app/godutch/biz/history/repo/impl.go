package repo

import (
	"time"

	"github.com/blackhorseya/godutch/internal/pkg/base/contextx"
	"github.com/blackhorseya/godutch/internal/pkg/entity/event"
	"github.com/blackhorseya/godutch/internal/pkg/entity/user"
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
       h.total      as total,
       user.id      as "payer.id",
       user.email   as "payer.email",
       user.name    as "payer.name",
       h.created_at as created_at
from spend_history h
         join users user on user.id = h.payer_id
where h.id = ?`
	err = i.rw.GetContext(timeout, &ret, stmt, id)
	if err != nil {
		return nil, err
	}

	var members []*user.Member
	stmt1 := `
select user.id      as id,
       user.email   as email,
       user.name    as name,
       detail.value as value
from spend_details detail
         join users user on detail.user_id = user.id
where detail.spend_id = ?`
	err = i.rw.SelectContext(timeout, &members, stmt1, id)
	if err != nil {
		return nil, err
	}

	ret.Members = members

	return &ret, nil
}

func (i *impl) List(ctx contextx.Contextx, actID int64, limit, offset int) (infos []*event.Record, err error) {
	timeout, cancel := contextx.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var ret []*event.Record
	stmt := `
select h.id       as id,
       h.remark   as remark,
       h.total    as total,
       user.id    as "payer.id",
       user.email as "payer.email",
       user.name  as "payer.name",
       h.created_at
from spend_history h
         join users user on h.payer_id = user.id
where h.activity_id = ?
order by h.created_at desc
limit ? offset ?`
	err = i.rw.SelectContext(timeout, &ret, stmt, actID, limit, offset)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (i *impl) Create(ctx contextx.Contextx, created *event.Record) (info *event.Record, err error) {
	timeout, cancel := contextx.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	stmt := `insert into spend_history (id, activity_id, payer_id, remark, total, created_at) 
values (:id, :activity.id, :payer.id, :remark, :total, :created_at)`
	_, err = i.rw.NamedExecContext(timeout, stmt, created)
	if err != nil {
		return nil, err
	}

	type detail struct {
		SpendID int64 `json:"spend_id" db:"spend_id"`
		UserID  int64 `json:"user_id" db:"user_id"`
		Value   int64 `json:"value" db:"value"`
	}
	var details []*detail
	for _, member := range created.Members {
		details = append(details, &detail{SpendID: created.ID, UserID: member.Id, Value: member.Value})
	}
	stmt1 := `insert into spend_details (spend_id, user_id, value) values (:spend_id, :user_id, :value)`
	_, err = i.rw.NamedExecContext(timeout, stmt1, details)
	if err != nil {
		return nil, err
	}

	return created, nil
}

func (i *impl) Delete(ctx contextx.Contextx, id int64) error {
	timeout, cancel := contextx.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	stmt := `delete from spend_history where id = ?`
	_, err := i.rw.ExecContext(timeout, stmt, id)
	if err != nil {
		return err
	}

	stmt1 := `delete from spend_details where spend_id = ?`
	_, err = i.rw.ExecContext(timeout, stmt1, id)
	if err != nil {
		return err
	}

	return nil
}
