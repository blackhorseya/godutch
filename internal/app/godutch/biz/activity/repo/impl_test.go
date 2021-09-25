package repo

import (
	"database/sql"
	"reflect"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/blackhorseya/godutch/internal/pkg/base/contextx"
	"github.com/blackhorseya/godutch/internal/pkg/entity/event"
	"github.com/blackhorseya/godutch/internal/pkg/entity/user"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

var (
	id1 = int64(1)

	userID1 = int64(1)

	user1 = &user.Profile{
		ID:    userID1,
		Email: "test",
		Name:  "test",
	}

	act1 = &event.Activity{
		ID:      id1,
		OwnerID: userID1,
		Name:    "test",
		Owner:   user1,
		Members: []*user.Profile{user1},
	}

	act2 = &event.Activity{
		ID:   id1,
		Name: "test",
	}
)

type repoSuite struct {
	suite.Suite
	mock sqlmock.Sqlmock
	repo IRepo
}

func (s *repoSuite) SetupTest() {
	logger, _ := zap.NewDevelopment()
	db, mock, _ := sqlmock.New()

	repo, err := CreateRepo(logger, sqlx.NewDb(db, "mysql"))
	if err != nil {
		panic(err)
	}

	s.repo = repo
	s.mock = mock
}

func TestRepoSuite(t *testing.T) {
	suite.Run(t, new(repoSuite))
}

func (s *repoSuite) Test_impl_GetByID() {
	stmt := "SELECT act.id, act.name, act.owner_id, owner.id \"owner.id\", owner.email \"owner.email\", owner.name \"owner.name\", act.created_at FROM activities act JOIN users owner ON owner.id = act.owner_id"

	type args struct {
		id     int64
		userID int64
		mock   func()
	}
	tests := []struct {
		name     string
		args     args
		wantInfo *event.Activity
		wantErr  bool
	}{
		{
			name: "get by id then error",
			args: args{id: id1, userID: userID1, mock: func() {
				s.mock.ExpectQuery(stmt).
					WithArgs(id1).
					WillReturnError(errors.New("error"))
			}},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "get by id then not found",
			args: args{id: id1, userID: userID1, mock: func() {
				s.mock.ExpectQuery(stmt).
					WithArgs(id1).
					WillReturnError(sql.ErrNoRows)
			}},
			wantInfo: nil,
			wantErr:  false,
		},
		{
			name: "get by id then success",
			args: args{id: id1, userID: userID1, mock: func() {
				s.mock.ExpectQuery(stmt).
					WithArgs(id1).
					WillReturnRows(sqlmock.NewRows([]string{"id", "name", "owner_id", "owner.id", "owner.email", "owner.name", "created_at"}).
						AddRow(act1.ID, act1.Name, act1.OwnerID, user1.ID, user1.Email, user1.Name, act1.CreatedAt))
			}},
			wantInfo: act1,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotInfo, err := s.repo.GetByID(contextx.Background(), tt.args.id, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotInfo, tt.wantInfo) {
				t.Errorf("GetByID() gotInfo = %v, want %v", gotInfo, tt.wantInfo)
			}
		})
	}
}

func (s *repoSuite) Test_impl_Create() {
	stmt := "INSERT INTO activities"
	stmt1 := "INSERT INTO activities_users_map"

	type args struct {
		created *event.Activity
		mock    func()
	}
	tests := []struct {
		name     string
		args     args
		wantInfo *event.Activity
		wantErr  bool
	}{
		{
			name: "create then error",
			args: args{created: act1, mock: func() {
				s.mock.ExpectExec(stmt).
					WithArgs(act1.ID, act1.Name, act1.OwnerID, act1.CreatedAt).
					WillReturnError(errors.New("error"))
			}},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "insert map then error",
			args: args{created: act1, mock: func() {
				s.mock.ExpectExec(stmt).
					WithArgs(act1.ID, act1.Name, act1.OwnerID, act1.CreatedAt).
					WillReturnResult(sqlmock.NewResult(1, 1))
				s.mock.ExpectExec(stmt1).
					WithArgs(act1.ID, user1.ID).
					WillReturnError(errors.New("error"))
			}},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "create then success",
			args: args{created: act1, mock: func() {
				s.mock.ExpectExec(stmt).
					WithArgs(act1.ID, act1.Name, act1.OwnerID, act1.CreatedAt).
					WillReturnResult(sqlmock.NewResult(1, 1))
				s.mock.ExpectExec(stmt1).
					WithArgs(act1.ID, user1.ID).
					WillReturnResult(sqlmock.NewResult(1, 1))
			}},
			wantInfo: act1,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotInfo, err := s.repo.Create(contextx.Background(), tt.args.created)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotInfo, tt.wantInfo) {
				t.Errorf("Create() gotInfo = %v, want %v", gotInfo, tt.wantInfo)
			}
		})
	}
}

func (s *repoSuite) Test_impl_List() {
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
JOIN users owner ON owner.id = act.owner_id`

	stmt1 := `
SELECT member.id    AS id,
       member.email AS email,
       member.name  AS name
FROM activities act
         JOIN activities_users_map map on act.id = map.activity_id
         JOIN users member on map.user_id = member.id`

	type args struct {
		userID int64
		limit  int
		offset int
		mock   func()
	}
	tests := []struct {
		name      string
		args      args
		wantInfos []*event.Activity
		wantErr   bool
	}{
		{
			name: "list then error",
			args: args{userID: userID1, limit: 5, offset: 0, mock: func() {
				s.mock.ExpectQuery(stmt).
					WithArgs(userID1, 5, 0).
					WillReturnError(errors.New("error"))
			}},
			wantInfos: nil,
			wantErr:   true,
		},
		{
			name: "list then not found",
			args: args{userID: userID1, limit: 5, offset: 0, mock: func() {
				s.mock.ExpectQuery(stmt).
					WithArgs(userID1, 5, 0).
					WillReturnError(sql.ErrNoRows)
			}},
			wantInfos: nil,
			wantErr:   false,
		},
		{
			name: "list then success",
			args: args{userID: userID1, limit: 5, offset: 0, mock: func() {
				s.mock.ExpectQuery(stmt).
					WithArgs(userID1, 5, 0).
					WillReturnRows(sqlmock.NewRows([]string{"id", "name", "created_at"}).
						AddRow(act2.ID, act2.Name, act2.CreatedAt))
				s.mock.ExpectQuery(stmt1).
					WithArgs(act1.ID).
					WillReturnError(errors.New("error"))
			}},
			wantInfos: []*event.Activity{act2},
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotInfos, err := s.repo.List(contextx.Background(), tt.args.userID, tt.args.limit, tt.args.offset)
			if (err != nil) != tt.wantErr {
				t.Errorf("List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotInfos, tt.wantInfos) {
				t.Errorf("List() gotInfos = %v, want %v", gotInfos, tt.wantInfos)
			}
		})
	}
}

func (s *repoSuite) Test_impl_Count() {
	stmt := `SELECT COUNT(id) "c" FROM activities`

	type args struct {
		userID int64
		mock   func()
	}
	tests := []struct {
		name      string
		args      args
		wantTotal int
		wantErr   bool
	}{
		{
			name: "count then error",
			args: args{userID: userID1, mock: func() {
				s.mock.ExpectQuery(regexp.QuoteMeta(stmt)).
					WithArgs(userID1).
					WillReturnError(errors.New("error"))
			}},
			wantTotal: 0,
			wantErr:   true,
		},
		{
			name: "count then success",
			args: args{userID: userID1, mock: func() {
				s.mock.ExpectQuery(regexp.QuoteMeta(stmt)).
					WithArgs(userID1).
					WillReturnRows(sqlmock.NewRows([]string{"c"}).
						AddRow(10))
			}},
			wantTotal: 10,
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotTotal, err := s.repo.Count(contextx.Background(), tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Count() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotTotal != tt.wantTotal {
				t.Errorf("Count() gotTotal = %v, want %v", gotTotal, tt.wantTotal)
			}
		})
	}
}

func (s *repoSuite) Test_impl_Update() {
	stmt := "UPDATE activities"

	type args struct {
		updated *event.Activity
		mock    func()
	}
	tests := []struct {
		name     string
		args     args
		wantInfo *event.Activity
		wantErr  bool
	}{
		{
			name: "update then error",
			args: args{updated: act1, mock: func() {
				s.mock.ExpectExec(stmt).WithArgs(act1.Name, act1.ID).WillReturnError(errors.New("error"))
			}},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "update then success",
			args: args{updated: act1, mock: func() {
				s.mock.ExpectExec(stmt).WithArgs(act1.Name, act1.ID).WillReturnResult(sqlmock.NewResult(1, 1))
			}},
			wantInfo: act1,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotInfo, err := s.repo.Update(contextx.Background(), tt.args.updated)
			if (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotInfo, tt.wantInfo) {
				t.Errorf("Update() gotInfo = %v, want %v", gotInfo, tt.wantInfo)
			}
		})
	}
}

func (s *repoSuite) Test_impl_Delete() {
	stmt := "DELETE FROM activities"

	type args struct {
		id     int64
		userID int64
		mock   func()
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "delete then error",
			args: args{id: id1, userID: userID1, mock: func() {
				s.mock.ExpectExec(stmt).WithArgs(id1, userID1).WillReturnError(errors.New("error"))
			}},
			wantErr: true,
		},
		{
			name: "delete then success",
			args: args{id: id1, userID: userID1, mock: func() {
				s.mock.ExpectExec(stmt).WithArgs(id1, userID1).WillReturnResult(sqlmock.NewResult(1, 1))
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			if err := s.repo.Delete(contextx.Background(), tt.args.id, tt.args.userID); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func (s *repoSuite) Test_impl_AddMembers() {
	stmt := `INSERT INTO activities_users_map`

	type args struct {
		id       int64
		newUsers []*user.Profile
		mock     func()
	}
	tests := []struct {
		name     string
		args     args
		wantInfo *event.Activity
		wantErr  bool
	}{
		{
			name: "add members then error",
			args: args{id: id1, newUsers: []*user.Profile{user1}, mock: func() {
				s.mock.ExpectExec(stmt).
					WithArgs(act1.ID, user1.ID).
					WillReturnError(errors.New("error"))
			}},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "add members then success",
			args: args{id: id1, newUsers: []*user.Profile{user1}, mock: func() {
				s.mock.ExpectExec(stmt).
					WithArgs(act1.ID, user1.ID).
					WillReturnResult(sqlmock.NewResult(1, 1))
			}},
			wantInfo: nil,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotInfo, err := s.repo.AddMembers(contextx.Background(), tt.args.id, tt.args.newUsers)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddMembers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotInfo, tt.wantInfo) {
				t.Errorf("AddMembers() gotInfo = %v, want %v", gotInfo, tt.wantInfo)
			}
		})
	}
}

func (s *repoSuite) Test_impl_GetByEmails() {
	stmt := "SELECT id, email, name FROM users"

	type args struct {
		emails []string
		mock   func()
	}
	tests := []struct {
		name      string
		args      args
		wantInfos []*user.Profile
		wantErr   bool
	}{
		{
			name: "get by email then error",
			args: args{emails: []string{"test"}, mock: func() {
				s.mock.ExpectQuery(stmt).
					WithArgs(user1.Email).
					WillReturnError(errors.New("error"))
			}},
			wantInfos: nil,
			wantErr:   false,
		},
		{
			name: "get by email then success",
			args: args{emails: []string{"test"}, mock: func() {
				s.mock.ExpectQuery(stmt).
					WithArgs(user1.Email).
					WillReturnRows(sqlmock.NewRows([]string{"id", "email", "name"}).
						AddRow(user1.ID, user1.Email, user1.Name))
			}},
			wantInfos: []*user.Profile{user1},
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotInfos, err := s.repo.GetByEmails(contextx.Background(), tt.args.emails)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByEmails() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotInfos, tt.wantInfos) {
				t.Errorf("GetByEmails() gotInfos = %v, want %v", gotInfos, tt.wantInfos)
			}
		})
	}
}
