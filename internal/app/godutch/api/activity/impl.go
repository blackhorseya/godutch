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
	// todo: 2021-09-25|01:03|Sean|impl me
	panic("implement me")
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
	// todo: 2021-09-25|01:03|Sean|impl me
	panic("implement me")
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
	// todo: 2021-09-25|01:03|Sean|impl me
	panic("implement me")
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
