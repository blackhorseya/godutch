package activity

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/blackhorseya/godutch/internal/app/godutch/biz/activity/mocks"
	"github.com/blackhorseya/godutch/internal/pkg/entity/er"
	"github.com/blackhorseya/godutch/internal/pkg/entity/event"
	"github.com/blackhorseya/godutch/internal/pkg/infra/transports/http/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

var (
	id1 = int64(1)

	act1 = &event.Activity{
		ID:   id1,
		Name: "test",
	}
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
	s.r.GET("/api/v1/activities/:id", s.handler.GetByID)

	type args struct {
		id   int64
		mock func()
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
	}{
		{
			name: "get by id then 500",
			args: args{id: id1, mock: func() {
				s.mock.On("GetByID", mock.Anything, id1).Return(nil, er.ErrGetActivityByID).Once()
			}},
			wantCode: 500,
		},
		{
			name: "get by id then 200",
			args: args{id: id1, mock: func() {
				s.mock.On("GetByID", mock.Anything, id1).Return(act1, nil).Once()
			}},
			wantCode: 200,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			uri := fmt.Sprintf("/api/v1/activities/%v", tt.args.id)
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
	s.r.GET("/api/v1/activities", s.handler.List)

	type args struct {
		page string
		size string
		mock func()
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
	}{
		{
			name:     "a 10 then 400",
			args:     args{page: "a", size: "10"},
			wantCode: 400,
		},
		{
			name:     "1 b then 400",
			args:     args{page: "1", size: "b"},
			wantCode: 400,
		},
		{
			name: "1 5 then 500",
			args: args{page: "1", size: "5", mock: func() {
				s.mock.On("List", mock.Anything, 1, 5).Return(nil, 0, er.ErrListActivities).Once()
			}},
			wantCode: 500,
		},
		{
			name: "1 5 then 200",
			args: args{page: "1", size: "5", mock: func() {
				s.mock.On("List", mock.Anything, 1, 5).Return([]*event.Activity{act1}, 10, nil).Once()
			}},
			wantCode: 200,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			uri := fmt.Sprintf("/api/v1/activities?page=%v&size=%v", tt.args.page, tt.args.size)
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

func (s *handlerSuite) Test_impl_NewWithMembers() {
	s.r.POST("/api/v1/activities", s.handler.NewWithMembers)

	type args struct {
		name   string
		emails []string
		mock   func()
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
	}{
		{
			name:     "missing name then 400",
			args:     args{name: "", emails: []string{"test"}},
			wantCode: 400,
		},
		{
			name: "new with members then 500",
			args: args{name: "test", emails: []string{"test"}, mock: func() {
				s.mock.On("NewWithMembers", mock.Anything, "test", []string{"test"}).Return(nil, er.ErrCreateActivity).Once()
			}},
			wantCode: 500,
		},
		{
			name: "new with members then 201",
			args: args{name: "test", emails: []string{"test"}, mock: func() {
				s.mock.On("NewWithMembers", mock.Anything, "test", []string{"test"}).Return(act1, nil).Once()
			}},
			wantCode: 201,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			uri := fmt.Sprintf("/api/v1/activities")
			data, _ := json.Marshal(&reqNew{Name: tt.args.name, Emails: tt.args.emails})
			req := httptest.NewRequest(http.MethodPost, uri, bytes.NewBuffer(data))
			w := httptest.NewRecorder()
			s.r.ServeHTTP(w, req)

			got := w.Result()
			defer got.Body.Close()

			s.EqualValuesf(tt.wantCode, got.StatusCode, "NewWithMembers() code = %v, wantCode = %v", got.StatusCode, tt.wantCode)

			s.TearDownTest()
		})
	}
}

func (s *handlerSuite) Test_impl_ChangeName() {
	s.r.PATCH("/api/v1/activities/:id/name", s.handler.ChangeName)

	type args struct {
		id   int64
		name string
		mock func()
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
	}{
		{
			name:     "missing name then 400",
			args:     args{id: id1, name: ""},
			wantCode: 400,
		},
		{
			name: "change name then 500",
			args: args{id: id1, name: "test", mock: func() {
				s.mock.On("ChangeName", mock.Anything, id1, "test").Return(nil, er.ErrUpdateActivity).Once()
			}},
			wantCode: 500,
		},
		{
			name: "change name then 200",
			args: args{id: id1, name: "test", mock: func() {
				s.mock.On("ChangeName", mock.Anything, id1, "test").Return(act1, nil).Once()
			}},
			wantCode: 200,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			uri := fmt.Sprintf("/api/v1/activities/%v/name", tt.args.id)
			data, _ := json.Marshal(&reqName{Name: tt.args.name})
			req := httptest.NewRequest(http.MethodPatch, uri, bytes.NewBuffer(data))
			w := httptest.NewRecorder()
			s.r.ServeHTTP(w, req)

			got := w.Result()
			defer got.Body.Close()

			s.EqualValuesf(tt.wantCode, got.StatusCode, "ChangeName() code = %v, wantCode = %v", got.StatusCode, tt.wantCode)

			s.TearDownTest()
		})
	}
}

func (s *handlerSuite) Test_impl_Delete() {
	s.r.DELETE("/api/v1/activities/:id", s.handler.Delete)

	type args struct {
		id   int64
		mock func()
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
	}{
		{
			name: "delete then 500",
			args: args{id: id1, mock: func() {
				s.mock.On("Delete", mock.Anything, id1).Return(er.ErrDeleteActivity).Once()
			}},
			wantCode: 500,
		},
		{
			name: "delete then 204",
			args: args{id: id1, mock: func() {
				s.mock.On("Delete", mock.Anything, id1).Return(nil).Once()
			}},
			wantCode: 204,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			uri := fmt.Sprintf("/api/v1/activities/%v", tt.args.id)
			req := httptest.NewRequest(http.MethodDelete, uri, nil)
			w := httptest.NewRecorder()
			s.r.ServeHTTP(w, req)

			got := w.Result()
			defer got.Body.Close()

			s.EqualValuesf(tt.wantCode, got.StatusCode, "Delete() code = %v, wantCode = %v", got.StatusCode, tt.wantCode)

			s.TearDownTest()
		})
	}
}
