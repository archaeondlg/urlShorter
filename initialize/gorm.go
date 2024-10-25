package initialize

import (
	"os"
	"project/global"
	"project/model/common"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func Gorm() *gorm.DB {
	global.DB = ConnectMysql()
	return global.DB
}

func ConnectMysql() *gorm.DB {
	m := global.Config.Mysql
	if m.Dbname == "" {
		return nil
	}
	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(), // DSN data source name
		DefaultStringSize:         191,     // string 类型字段的默认长度
		SkipInitializeWithVersion: false,   // 根据版本自动配置
	}
	ext := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   m.Prefix,
			SingularTable: m.Singular,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
		SkipDefaultTransaction:                   true, // gorm默认写操作都通过事务，关闭后可提高性能
		// PrepareStmt: true,  // 对所有语句预编译，提高性能
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), ext); err != nil {
		return nil
	} else {
		db.InstanceSet("gorm:table_options", "ENGINE="+m.Engine)
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}

func Migrate() {
	err := global.DB.AutoMigrate(
		common.User{},
		common.ShortUrl{},
		common.RedirectRecord{},
	)
	if err != nil {
		global.Log.Error("init Tables fail")
		os.Exit(0)
	}
	global.Log.Info("init tables success")
}
