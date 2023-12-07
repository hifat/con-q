package routeV1

import (
	"github.com/hifat/con-q-api/internal/app/handler"
	"github.com/hifat/con-q-api/internal/app/middleware"

	"github.com/gin-gonic/gin"
	_ "github.com/hifat/con-q-api/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Route struct {
	router     *gin.RouterGroup
	middleware middleware.Middleware
	handler    handler.Handler
}

func New(router *gin.RouterGroup, middleware middleware.Middleware, handler handler.Handler) *Route {
	return &Route{
		router,
		middleware,
		handler,
	}
}

// @title           ConQ API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @Security bearer
// @securityDefinitions.apikey bearer
// @in header
// @name Authorization

// @BasePath /v1
func (r *Route) Register() {
	v1 := r.router.Group("v1")
	v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.NewHandler(), ginSwagger.InstanceName("v1")))

	healtzHandler := r.handler.Healtz
	v1.GET("/healtz", healtzHandler.Get)

	authMiddleware := r.middleware.Auth

	authRoute := v1.Group("auth")
	authHandler := r.handler.Auth
	authRoute.POST("/register", authHandler.Register)
	authRoute.POST("/login", authHandler.Login)
	authRoute.POST("/refresh-token", authMiddleware.AuthGuard(), authHandler.RefreshToken)
}
