package api

import (
	// import swagger docs
	_ "github.com/blackhorseya/godutch/api/docs"
	"github.com/blackhorseya/godutch/internal/app/godutch/api/health"
	"github.com/blackhorseya/godutch/internal/pkg/infra/transports/http"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// CreateInitHandlerFn serve caller to create init handler
func CreateInitHandlerFn(healthH health.IHandler) http.InitHandlers {
	return func(r *gin.Engine) {
		api := r.Group("api")
		{
			api.GET("readiness", healthH.Readiness)
			api.GET("liveness", healthH.Liveness)

			// open any environments can access swagger
			api.GET("docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		}
	}
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(
	health.ProviderSet,
	CreateInitHandlerFn,
)
