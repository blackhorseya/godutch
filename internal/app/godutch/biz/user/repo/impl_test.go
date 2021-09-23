package repo

import (
	"database/sql"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/blackhorseya/godutch/internal/pkg/base/contextx"
	"github.com/blackhorseya/godutch/internal/pkg/entity/user"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

var (
	email1 = "email"

	pass1 = "$2a$04$W7XkHbwTrBUistouvflijuB2JOnYW4iEZEHVGgTX1bSERjRPZgZR."

	token1 = "token"

	name1 = "name"

	info1 = &user.Profile{
		ID:       1,
		Email:    email1,
		Password: pass1,
		Token:    token1,
		Name:     name1,
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

	if repo, err := CreateRepo(logger, sqlx.NewDb(db, "mysql")); err != nil {
		panic(err)
	} else {
		s.repo = repo
	}

	s.mock = mock
}

func TestRepoSuite(t *testing.T) {
	suite.Run(t, new(repoSuite))
}

func (s *repoSuite) Test_impl_GetByID() {
	type args struct {
		id   int64
		mock func()
	}
	tests := []struct {
		name     string
		args     args
		wantInfo *user.Profile
		wantErr  bool
	}{
		{
			name: "get by id then info",
			args: args{id: 1, mock: func() {
				s.mock.ExpectQuery("select id, email, password, token, name, created_at from users").
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"id", "email", "password", "token", "name", "created_at"}).
						AddRow(1, email1, pass1, token1, name1, 0))
			}},
			wantInfo: info1,
			wantErr:  false,
		},
		{
			name: "not found",
			args: args{id: 1, mock: func() {
				s.mock.ExpectQuery("select id, email, password, token, name, created_at from users").
					WithArgs(1).
					WillReturnError(sql.ErrNoRows)
			}},
			wantInfo: nil,
			wantErr:  false,
		},
		{
			name: "get by id then error",
			args: args{id: 1, mock: func() {
				s.mock.ExpectQuery("select id, email, password, token, name, created_at from users").
					WithArgs(1).
					WillReturnError(errors.New("error"))
			}},
			wantInfo: nil,
			wantErr:  true,
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

func (s *repoSuite) Test_impl_GetByToken() {
	type args struct {
		token string
		mock  func()
	}
	tests := []struct {
		name     string
		args     args
		wantInfo *user.Profile
		wantErr  bool
	}{
		{
			name: "get by token the error",
			args: args{token: token1, mock: func() {
				s.mock.ExpectQuery("select id, email, password, token, name, created_at from users where token = ?").
					WithArgs(token1).
					WillReturnError(errors.New("error"))
			}},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "get by token then not found",
			args: args{token: token1, mock: func() {
				s.mock.ExpectQuery("select id, email, password, token, name, created_at from users where token = ?").
					WithArgs(token1).
					WillReturnError(sql.ErrNoRows)
			}},
			wantInfo: nil,
			wantErr:  false,
		},
		{
			name: "get by token then  info",
			args: args{token: token1, mock: func() {
				s.mock.ExpectQuery("select id, email, password, token, name, created_at from users where token = ?").
					WithArgs(token1).
					WillReturnRows(sqlmock.NewRows([]string{"id", "email", "password", "token", "name", "created_at"}).
						AddRow(1, email1, pass1, token1, name1, 0))
			}},
			wantInfo: info1,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotInfo, err := s.repo.GetByToken(contextx.Background(), tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotInfo, tt.wantInfo) {
				t.Errorf("GetByToken() gotInfo = %v, want %v", gotInfo, tt.wantInfo)
			}
		})
	}
}

func (s *repoSuite) Test_impl_GetByEmail() {
	type args struct {
		email string
		mock  func()
	}
	tests := []struct {
		name     string
		args     args
		wantInfo *user.Profile
		wantErr  bool
	}{
		{
			name: "get by email the error",
			args: args{email: email1, mock: func() {
				s.mock.ExpectQuery("select id, email, password, token, name, created_at from users where email = ?").
					WithArgs(email1).
					WillReturnError(errors.New("error"))
			}},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "get by email then not found",
			args: args{email: email1, mock: func() {
				s.mock.ExpectQuery("select id, email, password, token, name, created_at from users where email = ?").
					WithArgs(email1).
					WillReturnError(sql.ErrNoRows)
			}},
			wantInfo: nil,
			wantErr:  false,
		},
		{
			name: "get by email then  info",
			args: args{email: email1, mock: func() {
				s.mock.ExpectQuery("select id, email, password, token, name, created_at from users where email = ?").
					WithArgs(email1).
					WillReturnRows(sqlmock.NewRows([]string{"id", "email", "password", "token", "name", "created_at"}).
						AddRow(1, email1, pass1, token1, name1, 0))
			}},
			wantInfo: info1,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotInfo, err := s.repo.GetByEmail(contextx.Background(), tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotInfo, tt.wantInfo) {
				t.Errorf("GetByEmail() gotInfo = %v, want %v", gotInfo, tt.wantInfo)
			}
		})
	}
}

func (s *repoSuite) Test_impl_Register() {
	type args struct {
		newUser *user.Profile
		mock    func()
	}
	tests := []struct {
		name     string
		args     args
		wantInfo *user.Profile
		wantErr  bool
	}{
		{
			name: "register then error",
			args: args{newUser: info1, mock: func() {
				s.mock.ExpectExec("insert into users (id, email, password, token, name, created_at) values (:id, :email, :password, :token, :name, :created_at)").
					WithArgs(info1).
					WillReturnError(errors.New("error"))
			}},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "register then error",
			args: args{newUser: info1, mock: func() {
				s.mock.ExpectExec("insert into users").
					WithArgs(info1).
					WillReturnResult(sqlmock.NewResult(1, 1))
			}},
			wantInfo: nil,
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotInfo, err := s.repo.Register(contextx.Background(), tt.args.newUser)
			if (err != nil) != tt.wantErr {
				t.Errorf("Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotInfo, tt.wantInfo) {
				t.Errorf("Register() gotInfo = %v, want %v", gotInfo, tt.wantInfo)
			}
		})
	}
}

func (s *repoSuite) Test_impl_UpdateToken() {
	type args struct {
		updated *user.Profile
		mock    func()
	}
	tests := []struct {
		name     string
		args     args
		wantInfo *user.Profile
		wantErr  bool
	}{
		{
			name: "update token then error",
			args: args{updated: info1, mock: func() {
				s.mock.ExpectExec("update users set token=:token where id=:id").
					WithArgs(info1).
					WillReturnError(errors.New("error"))
			}},
			wantInfo: nil,
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotInfo, err := s.repo.UpdateToken(contextx.Background(), tt.args.updated)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotInfo, tt.wantInfo) {
				t.Errorf("UpdateToken() gotInfo = %v, want %v", gotInfo, tt.wantInfo)
			}
		})
	}
}
