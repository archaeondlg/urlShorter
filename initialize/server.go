package initialize

import (
	"fmt"
	"project/global"

	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func Run() {
	Router := Routers()

	address := fmt.Sprintf(":%d", global.Config.System.Addr)
	s := initServer(address, Router)

	global.Log.Info("server run success on ", zap.String("address", address))

	fmt.Printf(`
	默认自动化文档地址:http://127.0.0.1%s/swagger/index.html
	默认前端文件运行地址:http://127.0.0.1:8080
    `, address)
	global.Log.Error(s.ListenAndServe().Error())
}
