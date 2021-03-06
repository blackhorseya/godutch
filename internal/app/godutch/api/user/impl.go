package user

import (
	"net/http"

	"github.com/blackhorseya/godutch/internal/app/godutch/biz/user"
	"github.com/blackhorseya/godutch/internal/pkg/base/contextx"
	"github.com/blackhorseya/godutch/internal/pkg/entity/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type impl struct {
	logger *zap.Logger
	biz    user.IBiz
}

// NewImpl serve caller to create an IHandler
func NewImpl(logger *zap.Logger, biz user.IBiz) IHandler {
	return &impl{
		logger: logger.With(zap.String("type", "UserHandler")),
		biz:    biz,
	}
}

// Me
// @Summary Get myself
// @Description Get myself
// @Tags Users
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Success 200 {object} response.Response{data=user.Profile}
// @Failure 400 {object} er.APPError
// @Failure 404 {object} er.APPError
// @Failure 500 {object} er.APPError
// @Router /v1/users/me [get]
func (i *impl) Me(c *gin.Context) {
	ctx := c.MustGet("ctx").(contextx.Contextx)

	info := ctx.Value("user")

	c.JSON(http.StatusOK, response.OK.WithData(info))
}

// Signup
// @Summary Signup
// @Description Signup
// @Tags Auth
// @Accept application/x-www-form-urlencoded
// @Produce application/json
// @Param email formData string true "email"
// @Param password formData string true "password"
// @Param name formData string true "name"
// @Success 201 {object} response.Response{data=user.Profile}
// @Failure 400 {object} er.APPError
// @Failure 500 {object} er.APPError
// @Router /v1/auth/signup [post]
func (i *impl) Signup(c *gin.Context) {
	ctx := c.MustGet("ctx").(contextx.Contextx)

	email := c.PostForm("email")
	password := c.PostForm("password")
	name := c.PostForm("name")

	ret, err := i.biz.Signup(ctx, email, password, name)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, response.OK.WithData(ret))
}

// Login
// @Summary Login
// @Description Login
// @Tags Auth
// @Accept application/x-www-form-urlencoded
// @Produce application/json
// @Param email formData string true "email"
// @Param password formData string true "password"
// @Success 201 {object} response.Response{data=user.Profile}
// @Failure 400 {object} er.APPError
// @Failure 500 {object} er.APPError
// @Router /v1/auth/login [post]
func (i *impl) Login(c *gin.Context) {
	ctx := c.MustGet("ctx").(contextx.Contextx)

	email := c.PostForm("email")
	password := c.PostForm("password")

	ret, err := i.biz.Login(ctx, email, password)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, response.OK.WithData(ret))
}

// Logout
// @Summary Logout
// @Description Logout
// @Tags Auth
// @Produce application/json
// @Security ApiKeyAuth
// @Success 200 {object} response.Response
// @Failure 500 {object} er.APPError
// @Router /v1/auth/logout [delete]
func (i *impl) Logout(c *gin.Context) {
	ctx := c.MustGet("ctx").(contextx.Contextx)

	err := i.biz.Logout(ctx)
	if err != nil {
		c.Error(err)
		return
	}

	c.Status(http.StatusOK)
}
