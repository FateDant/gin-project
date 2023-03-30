package bootstrap

import (
	"gin/app/models"
	"gin/global"
	"go.uber.org/zap"
	"gopkg.in/natefinch/lumberjack.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

// InitializeDB 初始化数据库连接
func InitializeDB() *gorm.DB {
	// 根据驱动配置进行初始化
	switch global.App.Config.Database.Driver {
	case "mysql":
		return initMySqlGorm()
	default:
		return initMySqlGorm()
	}
}

func initMySqlGorm() *gorm.DB {
	//获取数据库配置
	dbConfig := global.App.Config.Database

	if dbConfig.Database == "" {
		return nil
	}
	dsn := dbConfig.UserName + ":" + dbConfig.Password +
		"@tcp(" + dbConfig.Host + ":" + strconv.Itoa(dbConfig.Port) + ")/" +
		dbConfig.Database + "?charset=" + dbConfig.Charset + "&parseTime=True&loc=Local" // strconv.Itoa 数字转换成字符串
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,            // 禁用自动创建外键约束
		Logger:                                   getGormLogger(), // 使用自定义 Logger
	}); err != nil {
		global.App.Log.Error("mysql connect failed,err:", zap.Any("err", err)) //通过zap 写入错误信息
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(dbConfig.MaxIdleConns)
		sqlDB.SetMaxOpenConns(dbConfig.MaxOpenConns)
		initMySqlTables(db)
		return db
	}
}

func initMySqlTables(db *gorm.DB) {
	err := db.AutoMigrate(models.User{})
	if err != nil {
		global.App.Log.Error("migrate table failed", zap.Any("err", err))
		os.Exit(0)
	}
}

// 自定义 gorm Write
func getGormLogWriter() logger.Writer {
	var writer io.Writer

	// 是否启用日志文件
	if global.App.Config.Database.EnableFileLogWriter {
		// 自定义writer 写文件
		writer = &lumberjack.Logger{
			Filename:   global.App.Config.Log.RootDir + "/" + global.App.Config.Database.LogFilename,
			MaxSize:    global.App.Config.Log.MaxSize,
			MaxBackups: global.App.Config.Log.MaxBackups,
			MaxAge:     global.App.Config.Log.MaxAge,
			Compress:   global.App.Config.Log.Compress,
		}
	} else {
		//默认 writer 控制台输出
		writer = os.Stdout
	}
	return log.New(writer, "\r\n", log.LstdFlags)
}

// 切换默认 Logger 使用的 Writer
func getGormLogger() logger.Interface {
	var logModel logger.LogLevel

	//数据库日志级别
	switch global.App.Config.Database.LogMode {
	case "silent":
		logModel = logger.Silent
	case "error":
		logModel = logger.Error
	case "warn":
		logModel = logger.Warn
	case "info":
		logModel = logger.Info
	default:
		logModel = logger.Info
	}

	return logger.New(getGormLogWriter(), logger.Config{
		SlowThreshold:             200 * time.Millisecond,                          //慢sql阈值
		LogLevel:                  logModel,                                        //日志级别
		IgnoreRecordNotFoundError: false,                                           //忽略ErrRecordNotFound（记录未找到）错误
		Colorful:                  !global.App.Config.Database.EnableFileLogWriter, //使用彩色打印
	})
}
