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
	stmt := `
SELECT 
       act.id, 
       act.name, 
       act.owner_id,
       owner.id "owner.id",
       owner.email "owner.email", 
       owner.name "owner.name",
       act.created_at 
FROM activities act
JOIN users owner ON owner.id = act.owner_id
WHERE act.id = ?`
	err = i.rw.GetContext(timeout, &ret, stmt, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &ret, nil
}

func (i *impl) GetByEmails(ctx contextx.Contextx, emails []string) (infos []*user.Profile, err error) {
	timeout, cancel := contextx.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var users []*user.Profile
	for _, email := range emails {
		profile := user.Profile{}
		stmt := `SELECT id, email, name FROM users WHERE email = ?`
		err := i.rw.GetContext(timeout, &profile, stmt, email)
		if err != nil {
			continue
		}

		users = append(users, &profile)
	}

	return users, nil
}

func (i *impl) Create(ctx contextx.Contextx, created *event.Activity) (info *event.Activity, err error) {
	timeout, cancel := contextx.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	stmt := `INSERT INTO activities (id, name, owner_id, created_at) VALUES (:id, :name, :owner_id, :created_at)`
	_, err = i.rw.NamedExecContext(timeout, stmt, created)
	if err != nil {
		return nil, err
	}

	type mapping struct {
		ActivityID int64 `json:"activity_id" db:"activity_id"`
		UserID     int64 `json:"user_id" db:"user_id"`
	}
	var members []*mapping
	for _, member := range created.Members {
		members = append(members, &mapping{ActivityID: created.ID, UserID: member.ID})
	}
	stmt = `INSERT INTO activities_users_map (activity_id, user_id) VALUES (:activity_id, :user_id)`
	_, err = i.rw.NamedExecContext(timeout, stmt, members)
	if err != nil {
		return nil, err
	}

	return created, nil
}

func (i *impl) AddMembers(ctx contextx.Contextx, id int64, newUsers []*user.Profile) (info *event.Activity, err error) {
	timeout, cancel := contextx.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	type mapping struct {
		ActivityID int64 `json:"activity_id" db:"activity_id"`
		UserID     int64 `json:"user_id" db:"user_id"`
	}
	var members []*mapping
	for _, newUser := range newUsers {
		members = append(members, &mapping{ActivityID: id, UserID: newUser.ID})
	}

	stmt := `INSERT INTO activities_users_map (activity_id, user_id) VALUES (:activity_id, :user_id)`
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
SELECT 
       act.id AS id, 
       act.name AS name, 
       act.created_at AS created_at, 
       owner.id "owner.id",
       owner.email "owner.email", 
       owner.name "owner.name", 
       owner.created_at "owner.created_at" 
FROM activities act 
JOIN users owner ON owner.id = act.owner_id 
WHERE owner_id = ? limit ? offset ?`
	err = i.rw.SelectContext(timeout, &ret, stmt, userID, limit, offset)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return ret, nil
}

func (i *impl) Count(ctx contextx.Contextx, userID int64) (total int, err error) {
	timeout, cancel := contextx.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var ret int
	stmt := `SELECT COUNT(id) "c" FROM activities WHERE owner_id = ?`
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

	stmt := `DELETE FROM activities WHERE id = ? and owner_id = ?`
	_, err := i.rw.ExecContext(timeout, stmt, id, userID)
	if err != nil {
		return err
	}

	return nil
}
