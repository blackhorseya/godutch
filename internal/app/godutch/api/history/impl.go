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

// GetByID
// @Summary Get a spend record by id
// @Description Get a spend record by id
// @Tags History
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Param id path integer true "ID of activity"
// @Param record_id path integer true "ID of record"
// @Success 200 {object} response.Response{data=event.Record}
// @Failure 400 {object} er.APPError
// @Failure 404 {object} er.APPError
// @Failure 500 {object} er.APPError
// @Router /v1/activities/{id}/records/{record_id} [get]
func (i *impl) GetByID(c *gin.Context) {
	// todo: 2021-09-27|20:42|Sean|impl me
	panic("implement me")
}

// List
// @Summary List all records of activity
// @Description List all records of activity
// @Tags History
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Param id path integer true "ID of activity"
// @Param page query integer false "page" default(1)
// @Param size query integer false "size of page" default(5)
// @Success 200 {object} response.Response{data=[]event.Record}
// @Failure 400 {object} er.APPError
// @Failure 404 {object} er.APPError
// @Failure 500 {object} er.APPError
// @Router /v1/activities/{id}/records [get]
func (i *impl) List(c *gin.Context) {
	// todo: 2021-09-27|20:42|Sean|impl me
	panic("implement me")
}

type member struct {
	ID    int64 `json:"id" binding:"required"`
	Value int   `json:"value" binding:"required"`
}

type reqNew struct {
	PayerID int64     `json:"payer_id" binding:"required"`
	Remark  string    `json:"remark" binding:"required"`
	Members []*member `json:"members" binding:"required"`
	Total   int       `json:"total" binding:"required"`
}

// NewRecord
// @Summary Create a record
// @Description Create a record
// @Tags History
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Param id path integer true "ID of activity"
// @Param created body reqNew true "created record"
// @Success 201 {object} response.Response{data=event.Record}
// @Failure 400 {object} er.APPError
// @Failure 500 {object} er.APPError
// @Router /v1/activities/{id}/records [post]
func (i *impl) NewRecord(c *gin.Context) {
	// todo: 2021-09-27|20:42|Sean|impl me
	panic("implement me")
}

// Delete
// @Summary Remove a record by id
// @Description Remove a record by id
// @Tags History
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Param id path integer true "ID of activity"
// @Param record_id path integer true "ID of record"
// @Success 204 {object} string
// @Failure 400 {object} er.APPError
// @Failure 500 {object} er.APPError
// @Router /v1/activities/{id}/records/{record_id} [delete]
func (i *impl) Delete(c *gin.Context) {
	// todo: 2021-09-27|20:42|Sean|impl me
	panic("implement me")
}
