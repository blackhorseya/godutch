package api

import (
	// import swagger docs
	_ "github.com/blackhorseya/godutch/api/docs"
	"github.com/blackhorseya/godutch/internal/pkg/infra/transports/http"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// CreateInitHandlerFn serve caller to create init handler
func CreateInitHandlerFn() http.InitHandlers {
	return func(r *gin.Engine) {
		api := r.Group("api")
		{
			api.GET("readiness", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "ok"})
			})
			api.GET("liveness", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "ok"})
			})

			// open any environments can access swagger
			api.GET("docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		}
	}
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(
	CreateInitHandlerFn,
)
