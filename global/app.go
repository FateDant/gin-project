package global

import (
	"gin/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var App = new(Application)

// Application  用来存放一些项目启动时的变量，便于调用
type Application struct {
	ConfigViper *viper.Viper
	Config      config.Configuration
	Log         *zap.Logger
	DB          *gorm.DB
}
