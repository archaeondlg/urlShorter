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
	authRouter := Router.Group("auth")
	authRouter.Use(middleware.JWT())
	{
		authRouter.POST("changePassword", api.AdminApi.ChangePasswd)
	}
}
