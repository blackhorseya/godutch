package repo

import (
	"database/sql"
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
		logger: logger.With(zap.String("type", "ActivityRepo")),
		rw:     rw,
	}
}

func (i *impl) GetByID(ctx contextx.Contextx, id, userID int64) (info *event.Activity, err error) {
	timeout, cancel := contextx.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	ret := event.Activity{}
	stmt := `SELECT id, name, created_at FROM activities WHERE id = ?`
	err = i.rw.GetContext(timeout, &ret, stmt, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &ret, nil
}

func (i *impl) Create(ctx contextx.Contextx, created *event.Activity) (info *event.Activity, err error) {
	// todo: 2021-09-23|22:44|Sean|impl me
	panic("implement me")
}

func (i *impl) List(ctx contextx.Contextx, id, userID int64, limit, offset int) (infos []*event.Activity, err error) {
	// todo: 2021-09-23|22:44|Sean|impl me
	panic("implement me")
}

func (i *impl) Count(ctx contextx.Contextx, userID int64) (total int, err error) {
	// todo: 2021-09-23|22:44|Sean|impl me
	panic("implement me")
}

func (i *impl) Update(ctx contextx.Contextx, updated *event.Activity) (info *event.Activity, err error) {
	// todo: 2021-09-23|22:44|Sean|impl me
	panic("implement me")
}

func (i *impl) Delete(ctx contextx.Contextx, id, userID int64) error {
	// todo: 2021-09-23|22:44|Sean|impl me
	panic("implement me")
}
