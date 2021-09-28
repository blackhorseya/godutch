package history

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/blackhorseya/godutch/internal/app/godutch/biz/history/mocks"
	"github.com/blackhorseya/godutch/internal/pkg/entity/er"
	"github.com/blackhorseya/godutch/internal/pkg/entity/event"
	"github.com/blackhorseya/godutch/internal/pkg/entity/user"
	"github.com/blackhorseya/godutch/internal/pkg/infra/transports/http/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

var (
	actID1 = int64(1)

	recordID1 = int64(1)

	userID1 = int64(1)

	userID2 = int64(2)

	record1 = &event.Record{ID: recordID1, Activity: &event.Activity{ID: actID1}}

	member1 = &member{ID: userID1, Value: 500}

	member2 = &member{ID: userID2, Value: 100}
)

type handlerSuite struct {
	suite.Suite
	r       *gin.Engine
	mock    *mocks.IBiz
	handler IHandler
}

func (s *handlerSuite) SetupTest() {
	logger := zap.NewNop()

	gin.SetMode(gin.TestMode)
	s.r = gin.New()
	s.r.Use(middlewares.ContextMiddleware())
	s.r.Use(middlewares.ResponseMiddleware())

	s.mock = new(mocks.IBiz)
	handler, err := CreateIHandler(logger, s.mock)
	if err != nil {
		panic(err)
	}

	s.handler = handler
}

func (s *handlerSuite) TearDownTest() {
	s.mock.AssertExpectations(s.T())
}

func TestHandlerSuite(t *testing.T) {
	suite.Run(t, new(handlerSuite))
}

func (s *handlerSuite) Test_impl_GetByID() {
	s.r.GET("/api/v1/activities/:id/records/:record_id", s.handler.GetByID)

	type args struct {
		actID    int64
		recordID int64
		mock     func()
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
	}{
		{
			name:     "missing act id then error",
			args:     args{recordID: recordID1},
			wantCode: 400,
		},
		{
			name:     "missing record id then error",
			args:     args{actID: actID1},
			wantCode: 400,
		},
		{
			name: "get by id then error",
			args: args{actID: actID1, recordID: recordID1, mock: func() {
				s.mock.On("GetByID", mock.Anything, recordID1).Return(nil, er.ErrGetRecordByID).Once()
			}},
			wantCode: 500,
		},
		{
			name: "get by id then success",
			args: args{actID: actID1, recordID: recordID1, mock: func() {
				s.mock.On("GetByID", mock.Anything, recordID1).Return(record1, nil).Once()
			}},
			wantCode: 200,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			uri := fmt.Sprintf("/api/v1/activities/%v/records/%v", tt.args.actID, tt.args.recordID)
			req := httptest.NewRequest(http.MethodGet, uri, nil)
			w := httptest.NewRecorder()
			s.r.ServeHTTP(w, req)

			got := w.Result()
			defer got.Body.Close()

			s.EqualValuesf(tt.wantCode, got.StatusCode, "GetByID() code = %v, wantCode = %v", got.StatusCode, tt.wantCode)

			s.TearDownTest()
		})
	}
}

func (s *handlerSuite) Test_impl_List() {
	s.r.GET("/api/v1/activities/:id/records", s.handler.List)

	type args struct {
		actID int64
		page  string
		size  string
		mock  func()
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
	}{
		{
			name:     "missing act id then error",
			args:     args{page: "1", size: "10"},
			wantCode: 400,
		},
		{
			name:     "invalid page then error",
			args:     args{actID: actID1, page: "a", size: "10"},
			wantCode: 400,
		},
		{
			name:     "invalid size then error",
			args:     args{actID: actID1, page: "1", size: "b"},
			wantCode: 400,
		},
		{
			name: "list then error",
			args: args{actID: actID1, page: "1", size: "10", mock: func() {
				s.mock.On("List", mock.Anything, actID1, 1, 10).Return(nil, er.ErrListRecords).Once()
			}},
			wantCode: 500,
		},
		{
			name: "list then success",
			args: args{actID: actID1, page: "1", size: "10", mock: func() {
				s.mock.On("List", mock.Anything, actID1, 1, 10).Return([]*event.Record{record1}, nil).Once()
			}},
			wantCode: 200,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			uri := fmt.Sprintf("/api/v1/activities/%v/records?page=%v&size=%v", tt.args.actID, tt.args.page, tt.args.size)
			req := httptest.NewRequest(http.MethodGet, uri, nil)
			w := httptest.NewRecorder()
			s.r.ServeHTTP(w, req)

			got := w.Result()
			defer got.Body.Close()

			s.EqualValuesf(tt.wantCode, got.StatusCode, "List() code = %v, wantCode = %v", got.StatusCode, tt.wantCode)

			s.TearDownTest()
		})
	}
}

func (s *handlerSuite) Test_impl_NewRecord() {
	s.r.POST("/api/v1/activities/:id/records", s.handler.NewRecord)

	type args struct {
		actID   int64
		payerID int64
		remark  string
		members []*member
		total   int
		mock    func()
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
	}{
		{
			name:     "missing act id then error",
			args:     args{payerID: userID1, members: []*member{member1, member2}, remark: "test", total: 1000},
			wantCode: 400,
		},
		{
			name:     "missing remark then error",
			args:     args{actID: actID1, payerID: userID1, members: []*member{member1, member2}, remark: "", total: 1000},
			wantCode: 400,
		},
		{
			name: "new record then error",
			args: args{actID: actID1, payerID: userID1, members: []*member{member1, member2}, remark: "test", total: 1000, mock: func() {
				s.mock.On("NewRecord", mock.Anything, actID1, userID1, "test", []*user.Member{
					{Id: userID1, Value: int64(member1.Value)},
					{Id: userID2, Value: int64(member2.Value)},
				}, 1000).Return(nil, er.ErrNewRecord).Once()
			}},
			wantCode: 500,
		},
		{
			name: "new record then success",
			args: args{actID: actID1, payerID: userID1, members: []*member{member1, member2}, remark: "test", total: 1000, mock: func() {
				s.mock.On("NewRecord", mock.Anything, actID1, userID1, "test", []*user.Member{
					{Id: userID1, Value: int64(member1.Value)},
					{Id: userID2, Value: int64(member2.Value)},
				}, 1000).Return(record1, nil).Once()
			}},
			wantCode: 201,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			uri := fmt.Sprintf("/api/v1/activities/%v/records", tt.args.actID)
			data, _ := json.Marshal(&reqNew{
				PayerID: tt.args.payerID,
				Remark:  tt.args.remark,
				Members: tt.args.members,
				Total:   tt.args.total,
			})
			req := httptest.NewRequest(http.MethodPost, uri, bytes.NewBuffer(data))
			w := httptest.NewRecorder()
			s.r.ServeHTTP(w, req)

			got := w.Result()
			defer got.Body.Close()

			s.EqualValuesf(tt.wantCode, got.StatusCode, "NewRecord() code = %v, wantCode = %v", got.StatusCode, tt.wantCode)

			s.TearDownTest()
		})
	}
}
