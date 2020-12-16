package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // mysql 驱动
	"lirawx.cn/go-web/config"
)

var (
	// DB 数据库对象
	DB *gorm.DB
)

// InitMySQL 初始化链接数据库
func InitMySQL(cfg config.Config) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Mysql.User, cfg.Mysql.Password, cfg.Mysql.Host, cfg.Mysql.Port, cfg.Mysql.DB)
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return
	}
	return DB.DB().Ping()
}

// Close 关闭链接数据库
func Close() {
	DB.Close()
}
