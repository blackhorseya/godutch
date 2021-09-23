package repo

import (
	"database/sql"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/blackhorseya/godutch/internal/pkg/base/contextx"
	"github.com/blackhorseya/godutch/internal/pkg/entity/event"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

var (
	id1 = int64(1)

	userID1 = int64(1)

	act1 = &event.Activity{
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
				s.mock.ExpectQuery("SELECT id, name, created_at FROM activities").
					WithArgs(id1).
					WillReturnError(errors.New("error"))
			}},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "get by id then not found",
			args: args{id: id1, userID: userID1, mock: func() {
				s.mock.ExpectQuery("SELECT id, name, created_at FROM activities").
					WithArgs(id1).
					WillReturnError(sql.ErrNoRows)
			}},
			wantInfo: nil,
			wantErr:  false,
		},
		{
			name: "get by id then success",
			args: args{id: id1, userID: userID1, mock: func() {
				s.mock.ExpectQuery("SELECT id, name, created_at FROM activities").
					WithArgs(id1).
					WillReturnRows(sqlmock.NewRows([]string{"id", "name", "created_at"}).
						AddRow(act1.ID, act1.Name, act1.CreatedAt))
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
