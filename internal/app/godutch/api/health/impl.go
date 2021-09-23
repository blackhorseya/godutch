package health

import (
	"net/http"

	"github.com/blackhorseya/godutch/internal/app/godutch/biz/health"
	"github.com/blackhorseya/godutch/internal/pkg/base/contextx"
	"github.com/blackhorseya/godutch/internal/pkg/entity/er"
	"github.com/blackhorseya/godutch/internal/pkg/entity/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type impl struct {
	logger *zap.Logger
	biz    health.IBiz
}

// NewImpl serve caller to create an IHandler
func NewImpl(logger *zap.Logger, biz health.IBiz) IHandler {
	return &impl{
		logger: logger.With(zap.String("type", "HealthHandler")),
		biz:    biz,
	}
}

// Readiness
// @Summary Readiness
// @Description Show application was ready to start accepting traffic
// @Tags Health
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response
// @Failure 500 {object} er.APPError
// @Router /readiness [get]
func (i *impl) Readiness(c *gin.Context) {
	ctx := c.MustGet("ctx").(contextx.Contextx)

	err := i.biz.Readiness(ctx)
	if err != nil {
		i.logger.Error(er.ErrReadiness.Error(), zap.Error(err))
		c.Error(er.ErrReadiness)
		return
	}

	c.JSON(http.StatusOK, response.OK)
}

// Liveness
// @Summary Liveness
// @Description to know when to restart an application
// @Tags Health
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response
// @Failure 500 {object} er.APPError
// @Router /liveness [get]
func (i *impl) Liveness(c *gin.Context) {
	ctx := c.MustGet("ctx").(contextx.Contextx)

	err := i.biz.Liveness(ctx)
	if err != nil {
		i.logger.Error(er.ErrReadiness.Error(), zap.Error(err))
		c.Error(er.ErrLiveness)
		return
	}

	c.JSON(http.StatusOK, response.OK)
}
