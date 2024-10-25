package router

import (
	"github.com/gin-gonic/gin"
)

type PublicRouter struct{}

func (s PublicRouter) Register(Router *gin.Engine) {
	{
		Router.GET("/:code", api.UrlApi.Access)
	}
}
