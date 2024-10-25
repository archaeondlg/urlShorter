package router

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type justFilesFilesystem struct {
	fs http.FileSystem
}

func (fs justFilesFilesystem) Open(name string) (http.File, error) {
	f, err := fs.fs.Open(name)
	if err != nil {
		return nil, err
	}

	stat, err := f.Stat()
	if stat.IsDir() {
		return nil, os.ErrPermission
	}

	return f, nil
}

type FileRouter struct{}

func (s *FileRouter) Register(Router *gin.Engine) {
	{
		// 如果想要不使用nginx代理前端网页，可以修改 web/.env.production 下的
		// VUE_APP_BASE_API = /
		// VUE_APP_BASE_PATH = http://localhost
		// 然后执行打包命令 npm run build。在打开下面3行注释
		// Router.Static("/favicon.ico", "./dist/favicon.ico")
		// Router.Static("/assets", "./dist/assets")   // dist里面的静态资源
		// Router.StaticFile("/", "./dist/index.html") // 前端网页入口页面

		// 静态文件
		// Router.StaticFS(global.Config.Local.StorePath, justFilesFilesystem{http.Dir(global.Config.Local.StorePath)})
	}
}
