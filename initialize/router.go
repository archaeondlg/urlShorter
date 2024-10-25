package initialize

import (
	"github.com/gin-gonic/gin"
	"project/global"
	"project/router"
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
		router.RouterGroup.Register(Router)
	}
	global.Log.Info("router register success")
	return Router
}
