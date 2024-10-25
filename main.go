package main

import (
	"go.uber.org/zap"
	"project/global"
	"project/initialize"
)

func main() {
	global.Viper = initialize.Viper()
	global.Log = initialize.Zap()
	zap.ReplaceGlobals(global.Log)
	global.DB = initialize.Gorm()
	if global.DB != nil {
		initialize.Migrate()
		db, _ := global.DB.DB()
		defer db.Close()
	}
	global.Redis = initialize.Redis()
	global.Snowflake = initialize.Snowflake()
	initialize.Run()
}
