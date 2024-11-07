package router

import (
	"project/middleware"

	"github.com/gin-gonic/gin"
)

type TenantRouter struct{}

func (s *TenantRouter) Register(Router *gin.Engine) {
	{
		Router.POST("/tenant/login", api.TenantApi.Login)
	}
	authRouter := Router.Group("tenant")
	authRouter.Use(middleware.JWT())
	{
		authRouter.POST("changePassword", api.TenantApi.ChangePasswd)
	}
	urlRouter := Router.Group("url")
	urlRouter.Use(middleware.JWT())
	{
		urlRouter.POST("create", api.UrlApi.Create)
	}
}
