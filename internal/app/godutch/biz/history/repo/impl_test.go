package repo

import (
	"reflect"
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

	actID1 = int64(1)

	userID1 = int64(1)

	user1 = &user.Member{Id: userID1}

	record1 = &event.Record{
		ID:      id1,
		Payer:   user1,
		Members: []*user.Member{user1},
	}

	record2 = &event.Record{
		ID:      id1,
		Payer:   user1,
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
	stmt := `
select h.id         as id,
       h.remark     as remark,
       h.total      as total,
       user.id      as "payer.id",
       user.email   as "payer.email",
       user.name    as "payer.name",
       h.created_at as created_at
from spend_history h
         join users user on user.id = h.payer_id`
	stmt1 := `
select user.id      as id,
       user.email   as email,
       user.name    as name,
       detail.value as value
from spend_details detail
         join users user on detail.user_id = user.id`

	type args struct {
		id   int64
		mock func()
	}
	tests := []struct {
		name     string
		args     args
		wantInfo *event.Record
		wantErr  bool
	}{
		{
			name: "get by id then error",
			args: args{id: id1, mock: func() {
				s.mock.ExpectQuery(stmt).WithArgs(id1).WillReturnError(errors.New("error"))
			}},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "get members then error",
			args: args{id: id1, mock: func() {
				s.mock.ExpectQuery(stmt).WithArgs(id1).
					WillReturnRows(sqlmock.NewRows([]string{"id", "remark", "payer.id", "payer.email", "payer.name", "total", "created_at"}).
						AddRow(record1.ID, record1.Remark, record1.Payer.Id, record1.Payer.Email, record1.Payer.Name, record1.Total, record1.CreatedAt))
				s.mock.ExpectQuery(stmt1).WithArgs(id1).WillReturnError(errors.New("error"))
			}},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "get by id with members then success",
			args: args{id: id1, mock: func() {
				s.mock.ExpectQuery(stmt).WithArgs(id1).
					WillReturnRows(sqlmock.NewRows([]string{"id", "remark", "payer.id", "payer.email", "payer.name", "total", "created_at"}).
						AddRow(record1.ID, record1.Remark, record1.Payer.Id, record1.Payer.Email, record1.Payer.Name, record1.Total, record1.CreatedAt))
				s.mock.ExpectQuery(stmt1).WithArgs(id1).
					WillReturnRows(sqlmock.NewRows([]string{"id", "email", "name", "value"}).
						AddRow(user1.Id, user1.Email, user1.Name, user1.Value))
			}},
			wantInfo: record1,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotInfo, err := s.repo.GetByID(contextx.Background(), tt.args.id)
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

func (s *repoSuite) Test_impl_List() {
	stmt := `
select h.id       as id,
       h.remark   as remark,
       h.total    as total,
       user.id    as "payer.id",
       user.email as "payer.email",
       user.name  as "payer.name",
       h.created_at
from spend_history h
         join users user on h.payer_id = user.id`

	type args struct {
		actID  int64
		limit  int
		offset int
		mock   func()
	}
	tests := []struct {
		name      string
		args      args
		wantInfos []*event.Record
		wantErr   bool
	}{
		{
			name: "list return error",
			args: args{actID: actID1, limit: 10, offset: 0, mock: func() {
				s.mock.ExpectQuery(stmt).WithArgs(actID1, 10, 0).WillReturnError(errors.New("error"))
			}},
			wantInfos: nil,
			wantErr:   true,
		},
		{
			name: "list return success",
			args: args{actID: actID1, limit: 10, offset: 0, mock: func() {
				s.mock.ExpectQuery(stmt).WithArgs(actID1, 10, 0).
					WillReturnRows(sqlmock.NewRows([]string{"id", "remark", "total", "payer.id", "payer.email", "payer.name", "created_at"}).
						AddRow(record1.ID, record1.Remark, record1.Total, record1.Payer.Id, record1.Payer.Email, record1.Payer.Name, record1.CreatedAt))
			}},
			wantInfos: []*event.Record{record2},
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotInfos, err := s.repo.List(contextx.Background(), tt.args.actID, tt.args.limit, tt.args.offset)
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

func (s *repoSuite) Test_impl_Delete() {
	stmt := `delete from spend_history`
	stmt1 := `delete from spend_details`

	type args struct {
		id   int64
		mock func()
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "delete then success",
			args: args{id: id1, mock: func() {
				s.mock.ExpectExec(stmt).WithArgs(id1).WillReturnResult(sqlmock.NewResult(1, 1))
				s.mock.ExpectExec(stmt1).WithArgs(id1).WillReturnResult(sqlmock.NewResult(1, 1))
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			if err := s.repo.Delete(contextx.Background(), tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func (s *repoSuite) Test_impl_Create() {
	stmt := `insert into spend_history`
	stmt1 := `insert into spend_details`

	type args struct {
		created *event.Record
		mock    func()
	}
	tests := []struct {
		name     string
		args     args
		wantInfo *event.Record
		wantErr  bool
	}{
		{
			name: "insert spend history then error",
			args: args{created: record1, mock: func() {
				s.mock.ExpectExec(stmt).
					WithArgs(record1.ID, record1.Activity.ID, record1.Payer.Id, record1.Remark, record1.Total, record1.CreatedAt).
					WillReturnError(errors.New("error"))
			}},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "insert spend details then error",
			args: args{created: record1, mock: func() {
				s.mock.ExpectExec(stmt).
					WithArgs(record1.ID, record1.Activity.ID, record1.Payer.Id, record1.Remark, record1.Total, record1.CreatedAt).
					WillReturnResult(sqlmock.NewResult(1, 1))
				s.mock.ExpectExec(stmt1).
					WithArgs(record1.ID, record1.Members[0].Id, record1.Members[0].Value).
					WillReturnError(errors.New("error"))
			}},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "insert then success",
			args: args{created: record1, mock: func() {
				s.mock.ExpectExec(stmt).
					WithArgs(record1.ID, record1.Activity.ID, record1.Payer.Id, record1.Remark, record1.Total, record1.CreatedAt).
					WillReturnResult(sqlmock.NewResult(1, 1))
				s.mock.ExpectExec(stmt1).
					WithArgs(record1.ID, record1.Members[0].Id, record1.Members[0].Value).
					WillReturnResult(sqlmock.NewResult(1, 1))
			}},
			wantInfo: record1,
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
