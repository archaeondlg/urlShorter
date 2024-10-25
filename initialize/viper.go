package initialize

import (
	"fmt"
	// "path/filepath"
	"project/global"

	"github.com/fsnotify/fsnotify"
	// "github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Viper() *viper.Viper {
	var config string

	// if configEnv := os.Getenv(internal.ConfigEnv); configEnv == "" { // 判断 internal.ConfigEnv 常量存储的环境变量是否为空
	// 	switch gin.Mode() {
	// 	case gin.DebugMode:
	// 		config = internal.ConfigDefaultFile
	// 	case gin.ReleaseMode:
	// 		config = internal.ConfigReleaseFile
	// 	case gin.TestMode:
	// 		config = internal.ConfigTestFile
	// 	}
	// 	fmt.Printf("您正在使用gin模式的%s环境名称,config的路径为%s\n", gin.Mode(), config)
	// } else { // internal.ConfigEnv 常量存储的环境变量不为空 将值赋值于config
	// 	config = configEnv
	// 	fmt.Printf("您正在使用%s环境变量,config的路径为%s\n", internal.ConfigEnv, config)
	// }
	config = "config.yaml"
	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err = v.Unmarshal(&global.Config); err != nil {
			fmt.Println(err)
		}
	})
	if err = v.Unmarshal(&global.Config); err != nil {
		panic(err)
	}

	return v
}
