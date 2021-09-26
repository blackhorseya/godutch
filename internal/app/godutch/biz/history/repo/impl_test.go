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

	userID1 = int64(1)

	user1 = &user.Member{Id: userID1}

	record1 = &event.Record{ID: id1, Payer: user1}
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
       user.id      as "payer.id",
       user.email   as "payer.email",
       user.name    as "payer.name",
       detail.value as total,
       h.created_at as created_at
from spend_history h
         join users user on user.id = h.payer_id
         join spend_details detail on h.payer_id = detail.user_id`

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
			name: "get by id then success",
			args: args{id: id1, mock: func() {
				s.mock.ExpectQuery(stmt).WithArgs(id1).
					WillReturnRows(sqlmock.NewRows([]string{"id", "remark", "payer.id", "payer.email", "payer.name", "total", "created_at"}).
						AddRow(record1.ID, record1.Remark, record1.Payer.Id, record1.Payer.Email, record1.Payer.Name, record1.Total, record1.CreatedAt))
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
