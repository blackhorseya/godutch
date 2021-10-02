package repo

import (
	"database/sql"
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

	var members []*user.Member
	stmt1 := `
SELECT member.id    AS id,
       member.email AS email,
       member.name  AS name,
       map.kind
FROM activities act
         JOIN activities_users_map map on act.id = map.activity_id
         JOIN users member on map.user_id = member.id
WHERE act.id = ?`
	err = i.rw.SelectContext(timeout, &members, stmt1, id)
	if err != nil {
		return nil, err
	}

	ret.Members = members

	return &ret, nil
}

func (i *impl) GetByEmails(ctx contextx.Contextx, emails []string) (infos []*user.Member, err error) {
	timeout, cancel := contextx.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var members []*user.Member
	for _, email := range emails {
		member := user.Member{}
		stmt := `SELECT id, email, name FROM users WHERE email = ?`
		err := i.rw.GetContext(timeout, &member, stmt, email)
		if err != nil {
			continue
		}

		members = append(members, &member)
	}

	return members, nil
}

func (i *impl) Create(ctx contextx.Contextx, created *event.Activity) (info *event.Activity, err error) {
	timeout, cancel := contextx.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	stmt := `INSERT INTO activities (id, name, created_at) VALUES (:id, :name, :created_at)`
	_, err = i.rw.NamedExecContext(timeout, stmt, created)
	if err != nil {
		return nil, err
	}

	type mapping struct {
		ActivityID int64     `json:"activity_id" db:"activity_id"`
		UserID     int64     `json:"user_id" db:"user_id"`
		Kind       user.Kind `json:"kind" db:"kind"`
	}
	var members []*mapping
	for _, member := range created.Members {
		members = append(members, &mapping{ActivityID: created.ID, UserID: member.Id, Kind: member.Kind})
	}
	stmt = `INSERT INTO activities_users_map (activity_id, user_id, kind) VALUES (:activity_id, :user_id, :kind)`
	_, err = i.rw.NamedExecContext(timeout, stmt, members)
	if err != nil {
		return nil, err
	}

	return created, nil
}

func (i *impl) AddMembers(ctx contextx.Contextx, id int64, newUsers []*user.Member) (info *event.Activity, err error) {
	timeout, cancel := contextx.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	type mapping struct {
		ActivityID int64     `json:"activity_id" db:"activity_id"`
		UserID     int64     `json:"user_id" db:"user_id"`
		Kind       user.Kind `json:"kind" db:"kind"`
	}
	var members []*mapping
	for _, newUser := range newUsers {
		members = append(members, &mapping{ActivityID: id, UserID: newUser.Id, Kind: newUser.Kind})
	}

	stmt := `INSERT INTO activities_users_map (activity_id, user_id, kind) VALUES (:activity_id, :user_id, :kind)`
	_, err = i.rw.NamedExecContext(timeout, stmt, members)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (i *impl) List(ctx contextx.Contextx, userID int64, limit, offset int) (infos []*event.Activity, err error) {
	timeout, cancel := contextx.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var ret []*event.Activity
	stmt := `
select act.id, act.name, act.created_at from activities act
join activities_users_map map on act.id = map.activity_id
where map.user_id = ? limit ? offset ?`
	err = i.rw.SelectContext(timeout, &ret, stmt, userID, limit, offset)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	for _, activity := range ret {
		var members []*user.Member
		stmt1 := `
SELECT member.id    AS id,
       member.email AS email,
       member.name  AS name,
       map.kind
FROM activities act
         JOIN activities_users_map map on act.id = map.activity_id
         JOIN users member on map.user_id = member.id
WHERE act.id = ?`
		err := i.rw.SelectContext(timeout, &members, stmt1, activity.ID)
		if err != nil {
			continue
		}

		activity.Members = members
	}

	return ret, nil
}

func (i *impl) Count(ctx contextx.Contextx, userID int64) (total int, err error) {
	timeout, cancel := contextx.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var ret int
	stmt := `select count(act.id) from activities act
join activities_users_map map on act.id = map.activity_id
where map.user_id = ?`
	row := i.rw.QueryRowxContext(timeout, stmt, userID)
	err = row.Scan(&ret)
	if err != nil {
		return 0, err
	}

	return ret, nil
}

func (i *impl) Update(ctx contextx.Contextx, updated *event.Activity) (info *event.Activity, err error) {
	timeout, cancel := contextx.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	stmt := `UPDATE activities set name=:name WHERE id = :id`
	_, err = i.rw.NamedExecContext(timeout, stmt, updated)
	if err != nil {
		return nil, err
	}

	return updated, nil
}

func (i *impl) Delete(ctx contextx.Contextx, id, userID int64) error {
	timeout, cancel := contextx.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	stmt := `DELETE FROM activities WHERE id = ?`
	_, err := i.rw.ExecContext(timeout, stmt, id)
	if err != nil {
		return err
	}

	stmt = `DELETE FROM activities_users_map WHERE activity_id = ?`
	_, err = i.rw.ExecContext(timeout, stmt, id)
	if err != nil {
		return err
	}

	return nil
}
