package initialize

import (
	"project/global"
	"project/router"

	"github.com/gin-gonic/gin"
)

// 初始化总路由

func Routers() *gin.Engine {
	Router := gin.New()
	Router.Use(gin.Recovery())
	if gin.Mode() == gin.DebugMode {
		Router.Use(gin.Logger())
	}

	// 跨域，如需跨域可以打开下面的注释
	// Router.Use(middleware.Cors()) // 按照配置的规则放行跨域请求

	{
		router.RouterGroup.PublicRouter.Register(Router)
		router.RouterGroup.AdminRouter.Register(Router)
		router.RouterGroup.TenantRouter.Register(Router)
	}
	global.Log.Info("router register success")
	return Router
}
