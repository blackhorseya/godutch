package activity

import (
	"net/http"
	"strconv"

	"github.com/blackhorseya/godutch/internal/app/godutch/biz/activity"
	"github.com/blackhorseya/godutch/internal/pkg/base/contextx"
	"github.com/blackhorseya/godutch/internal/pkg/entity/er"
	"github.com/blackhorseya/godutch/internal/pkg/entity/response"
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

type reqID struct {
	ID int64 `uri:"id" binding:"required"`
}

// GetByID
// @Summary Get an activity by id
// @Description Get an activity by id
// @Tags Activities
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Param id path integer true "ID of activity"
// @Success 200 {object} response.Response{data=event.Activity}
// @Failure 400 {object} er.APPError
// @Failure 404 {object} er.APPError
// @Failure 500 {object} er.APPError
// @Router /v1/activities/{id} [get]
func (i *impl) GetByID(c *gin.Context) {
	ctx := c.MustGet("ctx").(contextx.Contextx)

	var req reqID
	if err := c.ShouldBindUri(&req); err != nil {
		i.logger.Error(er.ErrInvalidID.Error(), zap.Error(err))
		c.Error(er.ErrInvalidID)
		return
	}

	ret, err := i.biz.GetByID(ctx, req.ID)
	if err != nil {
		i.logger.Error(er.ErrGetActivityByID.Error(), zap.Error(err))
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response.OK.WithData(ret))
}

// List
// @Summary List all activities
// @Description List all activities
// @Tags Activities
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Param page query integer false "page" default(1)
// @Param size query integer false "size of page" default(5)
// @Success 200 {object} response.Response{data=[]event.Activity}
// @Failure 400 {object} er.APPError
// @Failure 404 {object} er.APPError
// @Failure 500 {object} er.APPError
// @Router /v1/activities [get]
func (i *impl) List(c *gin.Context) {
	ctx := c.MustGet("ctx").(contextx.Contextx)

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		i.logger.Error(er.ErrInvalidPage.Error(), zap.Error(err), zap.String("page", c.Query("page")))
		c.Error(er.ErrInvalidPage)
		return
	}

	size, err := strconv.Atoi(c.DefaultQuery("size", "5"))
	if err != nil {
		i.logger.Error(er.ErrInvalidSize.Error(), zap.Error(err), zap.String("size", c.Query("size")))
		c.Error(er.ErrInvalidSize)
		return
	}

	ret, total, err := i.biz.List(ctx, page, size)
	if err != nil {
		i.logger.Error(er.ErrListActivities.Error(), zap.Error(err), zap.Int("page", page), zap.Int("size", size))
		c.Error(err)
		return
	}

	c.Header("X-Total-Count", strconv.Itoa(total))
	c.JSON(http.StatusOK, response.OK.WithData(ret))
}

type reqNew struct {
	Name   string   `json:"name" binding:"required"`
	Emails []string `json:"emails"`
}

// NewWithMembers
// @Summary Create an activity with members email
// @Description Create an activity with members email
// @Tags Activities
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Param created body reqNew true "created activity"
// @Success 201 {object} response.Response{data=event.Activity}
// @Failure 400 {object} er.APPError
// @Failure 500 {object} er.APPError
// @Router /v1/activities [post]
func (i *impl) NewWithMembers(c *gin.Context) {
	ctx := c.MustGet("ctx").(contextx.Contextx)

	var data *reqNew
	if err := c.ShouldBindJSON(&data); err != nil {
		i.logger.Error(er.ErrEmptyName.Error(), zap.Error(err))
		c.Error(er.ErrEmptyName)
		return
	}

	ret, err := i.biz.NewWithMembers(ctx, data.Name, data.Emails)
	if err != nil {
		i.logger.Error(er.ErrCreateActivity.Error(), zap.Error(err), zap.Any("payload", data))
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, response.OK.WithData(ret))
}

type reqName struct {
	Name string `json:"name" binding:"required"`
}

// ChangeName
// @Summary Update an activity of name by id
// @Description Update an activity of name by id
// @Tags Activities
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Param id path integer true "ID of activity"
// @Param updated body reqName true "updated activity"
// @Success 200 {object} response.Response{data=event.Activity}
// @Failure 400 {object} er.APPError
// @Failure 500 {object} er.APPError
// @Router /v1/activities/{id}/name [patch]
func (i *impl) ChangeName(c *gin.Context) {
	// todo: 2021-09-25|01:03|Sean|impl me
	panic("implement me")
}

// Delete
// @Summary Remove an activity by id
// @Description Remove an activity by id
// @Tags Activities
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Param id path integer true "ID of activity"
// @Success 204 {object} string
// @Failure 400 {object} er.APPError
// @Failure 500 {object} er.APPError
// @Router /v1/activities/{id} [delete]
func (i *impl) Delete(c *gin.Context) {
	// todo: 2021-09-25|01:03|Sean|impl me
	panic("implement me")
}
