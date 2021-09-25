package activity

import (
	"reflect"
	"testing"

	"github.com/blackhorseya/godutch/internal/app/godutch/biz/activity/repo/mocks"
	"github.com/blackhorseya/godutch/internal/pkg/base/contextx"
	"github.com/blackhorseya/godutch/internal/pkg/entity/er"
	"github.com/blackhorseya/godutch/internal/pkg/entity/event"
	"github.com/blackhorseya/godutch/internal/pkg/entity/user"
	"github.com/bwmarrin/snowflake"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

var (
	id1 = int64(1)

	userID1 = int64(1)

	user1 = &user.Profile{ID: userID1}

	ctx1 = contextx.WithValue(contextx.Background(), "user", user1)

	act1 = &event.Activity{
		ID:      id1,
		Name:    "test",
		OwnerID: userID1,
		Owner:   user1,
	}
)

type bizSuite struct {
	suite.Suite
	mock *mocks.IRepo
	biz  IBiz
}

func (s *bizSuite) SetupTest() {
	logger, _ := zap.NewDevelopment()
	node, _ := snowflake.NewNode(1)

	s.mock = new(mocks.IRepo)
	biz, err := CreateIBiz(logger, s.mock, node)
	if err != nil {
		panic(err)
	}

	s.biz = biz
}

func (s *bizSuite) TearDownTest() {
	s.mock.AssertExpectations(s.T())
}

func TestBizSuite(t *testing.T) {
	suite.Run(t, new(bizSuite))
}

func (s *bizSuite) Test_impl_GetByID() {
	type args struct {
		ctx  contextx.Contextx
		id   int64
		mock func()
	}
	tests := []struct {
		name     string
		args     args
		wantInfo *event.Activity
		wantErr  bool
	}{
		{
			name:     "missing user info in ctx then error",
			args:     args{id: id1, ctx: contextx.Background()},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "get by id then error",
			args: args{id: id1, ctx: ctx1, mock: func() {
				s.mock.On("GetByID", mock.Anything, id1, userID1).
					Return(nil, er.ErrGetActivityByID).Once()
			}},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "get by id then not found",
			args: args{id: id1, ctx: ctx1, mock: func() {
				s.mock.On("GetByID", mock.Anything, id1, userID1).
					Return(nil, nil).Once()
			}},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "get by id then success",
			args: args{id: id1, ctx: ctx1, mock: func() {
				s.mock.On("GetByID", mock.Anything, id1, userID1).
					Return(act1, nil).Once()
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

			gotInfo, err := s.biz.GetByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotInfo, tt.wantInfo) {
				t.Errorf("GetByID() gotInfo = %v, want %v", gotInfo, tt.wantInfo)
			}

			s.TearDownTest()
		})
	}
}

func (s *bizSuite) Test_impl_List() {
	type args struct {
		ctx  contextx.Contextx
		page int
		size int
		mock func()
	}
	tests := []struct {
		name      string
		args      args
		wantInfos []*event.Activity
		wantTotal int
		wantErr   bool
	}{
		{
			name:      "missing user info in ctx then error",
			args:      args{page: 0, size: 10, ctx: contextx.Background()},
			wantInfos: nil,
			wantTotal: 0,
			wantErr:   true,
		},
		{
			name:      "invalid page then error",
			args:      args{page: -1, size: 10, ctx: ctx1},
			wantInfos: nil,
			wantTotal: 0,
			wantErr:   true,
		},
		{
			name:      "invalid size then error",
			args:      args{page: 1, size: -1, ctx: ctx1},
			wantInfos: nil,
			wantTotal: 0,
			wantErr:   true,
		},
		{
			name: "list then error",
			args: args{page: 1, size: 10, ctx: ctx1, mock: func() {
				s.mock.On("List", mock.Anything, userID1, 10, 0).
					Return(nil, errors.New("error")).Once()
			}},
			wantInfos: nil,
			wantTotal: 0,
			wantErr:   true,
		},
		{
			name: "list then not found",
			args: args{page: 1, size: 10, ctx: ctx1, mock: func() {
				s.mock.On("List", mock.Anything, userID1, 10, 0).
					Return(nil, nil).Once()
			}},
			wantInfos: nil,
			wantTotal: 0,
			wantErr:   true,
		},
		{
			name: "count then error",
			args: args{page: 1, size: 10, ctx: ctx1, mock: func() {
				s.mock.On("List", mock.Anything, userID1, 10, 0).
					Return([]*event.Activity{act1}, nil).Once()
				s.mock.On("Count", mock.Anything, userID1).
					Return(0, errors.New("error")).Once()
			}},
			wantInfos: nil,
			wantTotal: 0,
			wantErr:   true,
		},
		{
			name: "list and count then success",
			args: args{page: 1, size: 10, ctx: ctx1, mock: func() {
				s.mock.On("List", mock.Anything, userID1, 10, 0).
					Return([]*event.Activity{act1}, nil).Once()
				s.mock.On("Count", mock.Anything, userID1).
					Return(10, nil).Once()
			}},
			wantInfos: []*event.Activity{act1},
			wantTotal: 10,
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotInfos, gotTotal, err := s.biz.List(tt.args.ctx, tt.args.page, tt.args.size)
			if (err != nil) != tt.wantErr {
				t.Errorf("List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotInfos, tt.wantInfos) {
				t.Errorf("List() gotInfos = %v, want %v", gotInfos, tt.wantInfos)
			}
			if gotTotal != tt.wantTotal {
				t.Errorf("List() gotTotal = %v, want %v", gotTotal, tt.wantTotal)
			}

			s.TearDownTest()
		})
	}
}

func (s *bizSuite) Test_impl_ChangeName() {
	type args struct {
		ctx  contextx.Contextx
		id   int64
		name string
		mock func()
	}
	tests := []struct {
		name     string
		args     args
		wantInfo *event.Activity
		wantErr  bool
	}{
		{
			name:     "missing user info in ctx then error",
			args:     args{id: id1, name: "test", ctx: contextx.Background()},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name:     "missing name then error",
			args:     args{id: id1, name: "", ctx: ctx1},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "get by id then error",
			args: args{id: id1, name: "test", ctx: ctx1, mock: func() {
				s.mock.On("GetByID", mock.Anything, id1, userID1).Return(nil, errors.New("error")).Once()
			}},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "get by id then not found",
			args: args{id: id1, name: "test", ctx: ctx1, mock: func() {
				s.mock.On("GetByID", mock.Anything, id1, userID1).Return(nil, nil).Once()
			}},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "update then error",
			args: args{id: id1, name: "test", ctx: ctx1, mock: func() {
				s.mock.On("GetByID", mock.Anything, id1, userID1).Return(act1, nil).Once()
				s.mock.On("Update", mock.Anything, act1).Return(nil, errors.New("error")).Once()
			}},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "update then success",
			args: args{id: id1, name: "test", ctx: ctx1, mock: func() {
				s.mock.On("GetByID", mock.Anything, id1, userID1).Return(act1, nil).Once()
				s.mock.On("Update", mock.Anything, act1).Return(act1, nil).Once()
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

			gotInfo, err := s.biz.ChangeName(tt.args.ctx, tt.args.id, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("ChangeName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotInfo, tt.wantInfo) {
				t.Errorf("ChangeName() gotInfo = %v, want %v", gotInfo, tt.wantInfo)
			}

			s.TearDownTest()
		})
	}
}

func (s *bizSuite) Test_impl_Delete() {
	type args struct {
		ctx  contextx.Contextx
		id   int64
		mock func()
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "missing user info in ctx then error",
			args:    args{id: id1, ctx: contextx.Background()},
			wantErr: true,
		},
		{
			name: "delete then error",
			args: args{id: id1, ctx: ctx1, mock: func() {
				s.mock.On("Delete", mock.Anything, id1, userID1).Return(errors.New("error")).Once()
			}},
			wantErr: true,
		},
		{
			name: "delete then success",
			args: args{id: id1, ctx: ctx1, mock: func() {
				s.mock.On("Delete", mock.Anything, id1, userID1).Return(nil).Once()
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			if err := s.biz.Delete(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}

			s.TearDownTest()
		})
	}
}

func (s *bizSuite) Test_impl_NewWithMembers() {
	type args struct {
		ctx   contextx.Contextx
		name  string
		email []string
		mock  func()
	}
	tests := []struct {
		name     string
		args     args
		wantInfo *event.Activity
		wantErr  bool
	}{
		{
			name:     "missing user info in ctx then error",
			args:     args{name: "test", email: []string{"test"}, ctx: contextx.Background()},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name:     "missing emails then error",
			args:     args{name: "test", email: []string{}, ctx: ctx1},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name:     "missing name then error",
			args:     args{name: "", email: []string{"test"}, ctx: ctx1},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "get by emails then error",
			args: args{name: "test", email: []string{"test"}, ctx: ctx1, mock: func() {
				s.mock.On("GetByEmails", mock.Anything, []string{"test"}).Return(nil, errors.New("error")).Once()
			}},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "new an act then error",
			args: args{name: "test", email: []string{"test"}, ctx: ctx1, mock: func() {
				s.mock.On("GetByEmails", mock.Anything, []string{"test"}).Return([]*user.Profile{user1}, nil).Once()
				s.mock.On("Create", mock.Anything, mock.Anything).Return(nil, errors.New("error")).Once()
			}},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "new an act then success",
			args: args{name: "test", email: []string{"test"}, ctx: ctx1, mock: func() {
				s.mock.On("GetByEmails", mock.Anything, []string{"test"}).Return([]*user.Profile{user1}, nil).Once()
				s.mock.On("Create", mock.Anything, mock.Anything).Return(act1, nil).Once()
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

			gotInfo, err := s.biz.NewWithMembers(tt.args.ctx, tt.args.name, tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewWithMembers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotInfo, tt.wantInfo) {
				t.Errorf("NewWithMembers() gotInfo = %v, want %v", gotInfo, tt.wantInfo)
			}

			s.TearDownTest()
		})
	}
}

func (s *bizSuite) Test_impl_InviteMembers() {
	type args struct {
		ctx    contextx.Contextx
		id     int64
		emails []string
		mock   func()
	}
	tests := []struct {
		name     string
		args     args
		wantInfo *event.Activity
		wantErr  bool
	}{
		{
			name: "get user by emails then error",
			args: args{id: id1, emails: []string{"test"}, ctx: ctx1, mock: func() {
				s.mock.On("GetByEmails", mock.Anything, []string{"test"}).Return(nil, errors.New("error")).Once()
			}},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "invite members then error",
			args: args{id: id1, emails: []string{"test"}, ctx: ctx1, mock: func() {
				s.mock.On("GetByEmails", mock.Anything, []string{"test"}).Return([]*user.Profile{user1}, nil).Once()
				s.mock.On("AddMembers", mock.Anything, id1, []*user.Profile{user1}).Return(nil, errors.New("error")).Once()
			}},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "invite members then success",
			args: args{id: id1, emails: []string{"test"}, ctx: ctx1, mock: func() {
				s.mock.On("GetByEmails", mock.Anything, []string{"test"}).Return([]*user.Profile{user1}, nil).Once()
				s.mock.On("AddMembers", mock.Anything, id1, []*user.Profile{user1}).Return(act1, nil).Once()
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

			gotInfo, err := s.biz.InviteMembers(tt.args.ctx, tt.args.id, tt.args.emails)
			if (err != nil) != tt.wantErr {
				t.Errorf("InviteMembers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotInfo, tt.wantInfo) {
				t.Errorf("InviteMembers() gotInfo = %v, want %v", gotInfo, tt.wantInfo)
			}

			s.TearDownTest()
		})
	}
}
