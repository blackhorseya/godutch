package history

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/blackhorseya/godutch/internal/app/godutch/biz/history/mocks"
	"github.com/blackhorseya/godutch/internal/pkg/entity/er"
	"github.com/blackhorseya/godutch/internal/pkg/entity/event"
	"github.com/blackhorseya/godutch/internal/pkg/infra/transports/http/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

var (
	actID1 = int64(1)

	recordID1 = int64(1)

	record1 = &event.Record{ID: recordID1, Activity: &event.Activity{ID: actID1}}
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
