package activity

import (
	"testing"

	"github.com/blackhorseya/godutch/internal/app/godutch/biz/activity/repo/mocks"
	"github.com/bwmarrin/snowflake"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
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
