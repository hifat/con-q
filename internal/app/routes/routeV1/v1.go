package routeV1

import (
	"github.com/hifat/con-q/internal/app/handler"

	"github.com/gin-gonic/gin"
)

type Route struct {
	router  *gin.RouterGroup
	handler handler.Handler
}

func New(router *gin.RouterGroup, handler handler.Handler) *Route {
	return &Route{
		router,
		handler,
	}
}

func (r *Route) Register() {
	v1 := r.router.Group("v1")

	healtzHandler := r.handler.Healtz
	v1.GET("/healtz", healtzHandler.Get)
}
