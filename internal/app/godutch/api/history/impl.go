package history

import (
	"github.com/blackhorseya/godutch/internal/app/godutch/biz/history"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type impl struct {
	logger *zap.Logger
	biz    history.IBiz
}

// NewImpl serve caller to create an IHandler
func NewImpl(logger *zap.Logger, biz history.IBiz) IHandler {
	return &impl{
		logger: logger.With(zap.String("type", "HistoryHandler")),
		biz:    biz,
	}
}

func (i *impl) GetByID(c *gin.Context) {
	// todo: 2021-09-27|20:42|Sean|impl me
	panic("implement me")
}

func (i *impl) List(c *gin.Context) {
	// todo: 2021-09-27|20:42|Sean|impl me
	panic("implement me")
}

func (i *impl) NewRecord(c *gin.Context) {
	// todo: 2021-09-27|20:42|Sean|impl me
	panic("implement me")
}

func (i *impl) Delete(c *gin.Context) {
	// todo: 2021-09-27|20:42|Sean|impl me
	panic("implement me")
}
