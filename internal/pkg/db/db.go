package db

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"go-template/internal/conf"

	logdef "go-template/internal/pkg/log"
)

type Config struct {
	DSN             string // write data source name.
	MaxOpenConn     int    // open pool
	MaxIdleConn     int    // idle pool
	ConnMaxLifeTime int
}

var DB *gorm.DB

// InitDbConfig 初始化Db
func InitDbConfig(c *conf.Data) {
	log.Info("Initializing Mysql")
	var err error
	dsn := c.Database.Dsn
	maxIdleConns := c.Database.MaxIdleConn
	maxOpenConns := c.Database.MaxOpenConn
	connMaxLifetime := c.Database.ConnMaxLifeTime
	if DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		QueryFields: true,
		NamingStrategy: schema.NamingStrategy{
			//TablePrefix:   "",   // 表名前缀
			SingularTable: true, // 使用单数表名
		},
		Logger: logdef.NewGormLogger(),
	}); err != nil {
		panic(fmt.Errorf("初始化数据库失败: %s \n", err))
	}
	sqlDB, err := DB.DB()
	if sqlDB != nil {
		sqlDB.SetMaxIdleConns(int(maxIdleConns))                               // 空闲连接数
		sqlDB.SetMaxOpenConns(int(maxOpenConns))                               // 最大连接数
		sqlDB.SetConnMaxLifetime(time.Second * time.Duration(connMaxLifetime)) // 单位：秒
	}
	log.Info("Mysql: initialization completed")
}
