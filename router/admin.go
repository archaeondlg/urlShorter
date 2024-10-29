package router

import (
	"project/middleware"

	"github.com/gin-gonic/gin"
)

type AdminRouter struct{}

func (s AdminRouter) Register(Router *gin.Engine) {
	{
		Router.POST("/auth/login", api.AdminApi.Login)
	}
	Router.Use(middleware.JWT())
	authRouter := Router.Group("auth")
	{
		authRouter.POST("changePassword", api.AdminApi.ChangePasswd)
	}
	urlRouter := Router.Group("url")
	{
		urlRouter.POST("create", api.UrlApi.Create)
	}
}
