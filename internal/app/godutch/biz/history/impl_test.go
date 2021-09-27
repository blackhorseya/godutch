package history

import (
	"reflect"
	"testing"

	"github.com/blackhorseya/godutch/internal/app/godutch/biz/history/repo/mocks"
	"github.com/blackhorseya/godutch/internal/pkg/base/contextx"
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

	member1 = &user.Member{Id: id1}

	ctx1 = contextx.WithValue(contextx.Background(), "user", user1)

	record1 = &event.Record{ID: id1}
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
		wantInfo *event.Record
		wantErr  bool
	}{
		{
			name: "get by id then error",
			args: args{ctx: contextx.Background(), id: id1, mock: func() {
				s.mock.On("GetByID", mock.Anything, id1).Return(nil, errors.New("error")).Once()
			}},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "get by id then not found",
			args: args{ctx: contextx.Background(), id: id1, mock: func() {
				s.mock.On("GetByID", mock.Anything, id1).Return(nil, nil).Once()
			}},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "get by id then success",
			args: args{ctx: contextx.Background(), id: id1, mock: func() {
				s.mock.On("GetByID", mock.Anything, id1).Return(record1, nil).Once()
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

			gotInfo, err := s.biz.GetByID(tt.args.ctx, tt.args.id)
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

func (s *bizSuite) Test_impl_List() {
	type args struct {
		ctx   contextx.Contextx
		actID int64
		page  int
		size  int
		mock  func()
	}
	tests := []struct {
		name      string
		args      args
		wantInfos []*event.Record
		wantErr   bool
	}{
		{
			name:      "0 10 then error",
			args:      args{ctx: contextx.Background(), actID: id1, page: -1, size: 10},
			wantInfos: nil,
			wantErr:   true,
		},
		{
			name:      "1 0 then error",
			args:      args{ctx: contextx.Background(), actID: id1, page: 1, size: 0},
			wantInfos: nil,
			wantErr:   true,
		},
		{
			name: "list then error",
			args: args{ctx: contextx.Background(), actID: id1, page: 1, size: 10, mock: func() {
				s.mock.On("List", mock.Anything, id1, 10, 0).Return(nil, errors.New("error")).Once()
			}},
			wantInfos: nil,
			wantErr:   true,
		},
		{
			name: "list then not found",
			args: args{ctx: contextx.Background(), actID: id1, page: 1, size: 10, mock: func() {
				s.mock.On("List", mock.Anything, id1, 10, 0).Return(nil, nil).Once()
			}},
			wantInfos: nil,
			wantErr:   true,
		},
		{
			name: "list then success",
			args: args{ctx: contextx.Background(), actID: id1, page: 1, size: 10, mock: func() {
				s.mock.On("List", mock.Anything, id1, 10, 0).Return([]*event.Record{record1}, nil).Once()
			}},
			wantInfos: []*event.Record{record1},
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotInfos, err := s.biz.List(tt.args.ctx, tt.args.actID, tt.args.page, tt.args.size)
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
