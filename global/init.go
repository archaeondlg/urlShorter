package global

import (
	"project/config"

	"github.com/bwmarrin/snowflake"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Viper     *viper.Viper
	Config    config.Config
	Log       *zap.Logger
	DB        *gorm.DB
	Redis     redis.UniversalClient
	Snowflake *snowflake.Node
)
