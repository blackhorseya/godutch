package history

import (
	"net/http"
	"strconv"

	"github.com/blackhorseya/godutch/internal/app/godutch/biz/history"
	"github.com/blackhorseya/godutch/internal/pkg/base/contextx"
	"github.com/blackhorseya/godutch/internal/pkg/entity/er"
	"github.com/blackhorseya/godutch/internal/pkg/entity/response"
	"github.com/blackhorseya/godutch/internal/pkg/entity/user"
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

type reqRecordID struct {
	ID       int64 `uri:"id" binding:"required"`
	RecordID int64 `uri:"record_id" binding:"required"`
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
	ctx := c.MustGet("ctx").(contextx.Contextx)

	var req reqRecordID
	if err := c.ShouldBindUri(&req); err != nil {
		i.logger.Error(er.ErrInvalidID.Error(), zap.Error(err))
		_ = c.Error(er.ErrInvalidID)
		return
	}

	ret, err := i.biz.GetByID(ctx, req.RecordID)
	if err != nil {
		i.logger.Error(er.ErrGetRecordByID.Error(), zap.Error(err))
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response.OK.WithData(ret))
}

type reqID struct {
	ID int64 `uri:"id" binding:"required"`
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
	ctx := c.MustGet("ctx").(contextx.Contextx)

	var req reqID
	if err := c.ShouldBindUri(&req); err != nil {
		i.logger.Error(er.ErrInvalidID.Error(), zap.Error(err))
		_ = c.Error(er.ErrInvalidID)
		return
	}

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		i.logger.Error(er.ErrInvalidPage.Error(), zap.Error(err), zap.String("page", c.Query("page")))
		_ = c.Error(er.ErrInvalidPage)
		return
	}

	size, err := strconv.Atoi(c.DefaultQuery("size", "5"))
	if err != nil {
		i.logger.Error(er.ErrInvalidSize.Error(), zap.Error(err), zap.String("size", c.Query("size")))
		_ = c.Error(er.ErrInvalidSize)
		return
	}

	ret, err := i.biz.List(ctx, req.ID, page, size)
	if err != nil {
		i.logger.Error(er.ErrListRecords.Error(), zap.Error(err))
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response.OK.WithData(ret))
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
	ctx := c.MustGet("ctx").(contextx.Contextx)

	var req reqID
	if err := c.ShouldBindUri(&req); err != nil {
		i.logger.Error(er.ErrInvalidID.Error(), zap.Error(err))
		_ = c.Error(er.ErrInvalidID)
		return
	}

	var data *reqNew
	if err := c.ShouldBindJSON(&data); err != nil {
		i.logger.Error(er.ErrEmptyName.Error(), zap.Error(err))
		_ = c.Error(er.ErrEmptyName)
		return
	}

	var members []*user.Member
	for _, m := range data.Members {
		members = append(members, &user.Member{Id: m.ID, Value: int64(m.Value)})
	}
	ret, err := i.biz.NewRecord(ctx, req.ID, data.PayerID, data.Remark, members, data.Total)
	if err != nil {
		i.logger.Error(er.ErrNewRecord.Error(), zap.Error(err), zap.Any("payload", data))
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, response.OK.WithData(ret))
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
	ctx := c.MustGet("ctx").(contextx.Contextx)

	var req reqRecordID
	if err := c.ShouldBindUri(&req); err != nil {
		i.logger.Error(er.ErrInvalidID.Error(), zap.Error(err))
		_ = c.Error(er.ErrInvalidID)
		return
	}

	err := i.biz.Delete(ctx, req.RecordID, req.ID)
	if err != nil {
		i.logger.Error(er.ErrDeleteRecord.Error(), zap.Error(err))
		_ = c.Error(err)
		return
	}

	c.Status(http.StatusNoContent)
}
