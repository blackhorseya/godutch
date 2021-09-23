package api

import (
	// import swagger docs
	_ "github.com/blackhorseya/godutch/api/docs"
	"github.com/blackhorseya/godutch/internal/app/godutch/api/health"
	"github.com/blackhorseya/godutch/internal/app/godutch/api/user"
	userB "github.com/blackhorseya/godutch/internal/app/godutch/biz/user"
	"github.com/blackhorseya/godutch/internal/pkg/infra/transports/http"
	"github.com/blackhorseya/godutch/internal/pkg/infra/transports/http/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// CreateInitHandlerFn serve caller to create init handler
func CreateInitHandlerFn(userBiz userB.IBiz,
	healthH health.IHandler,
	userH user.IHandler) http.InitHandlers {
	return func(r *gin.Engine) {
		api := r.Group("api")
		{
			api.GET("readiness", healthH.Readiness)
			api.GET("liveness", healthH.Liveness)

			// open any environments can access swagger
			api.GET("docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

			v1 := api.Group("v1")
			{
				authG := v1.Group("auth")
				{
					authG.POST("signup", userH.Signup)
					authG.POST("login", userH.Login)
					authG.DELETE("logout", middlewares.AuthMiddleware(userBiz), userH.Logout)
				}

				userG := v1.Group("users")
				{
					userG.GET("me", middlewares.AuthMiddleware(userBiz), userH.Me)
				}
			}
		}
	}
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(
	health.ProviderSet,
	user.ProviderSet,
	CreateInitHandlerFn,
)
