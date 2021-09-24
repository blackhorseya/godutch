package activity

import (
	"github.com/blackhorseya/godutch/internal/app/godutch/biz/activity"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type impl struct {
	logger *zap.Logger
	biz    activity.IBiz
}

// NewImpl serve caller to create an IHandler
func NewImpl(logger *zap.Logger, biz activity.IBiz) IHandler {
	return &impl{
		logger: logger.With(zap.String("type", "ActivityHandler")),
		biz:    biz,
	}
}

func (i *impl) GetByID(c *gin.Context) {
	// todo: 2021-09-25|01:03|Sean|impl me
	panic("implement me")
}

func (i *impl) List(c *gin.Context) {
	// todo: 2021-09-25|01:03|Sean|impl me
	panic("implement me")
}

func (i *impl) NewWithMembers(c *gin.Context) {
	// todo: 2021-09-25|01:03|Sean|impl me
	panic("implement me")
}

func (i *impl) ChangeName(c *gin.Context) {
	// todo: 2021-09-25|01:03|Sean|impl me
	panic("implement me")
}

func (i *impl) Delete(c *gin.Context) {
	// todo: 2021-09-25|01:03|Sean|impl me
	panic("implement me")
}
