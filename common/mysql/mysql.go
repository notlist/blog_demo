package mysql

import (
	"fmt"
	"goadmin/common/config"
	"goadmin/common/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const DRIVER = "mysql"

var SqlSession *gorm.DB

func InitMySql() (err error) {
	SqlSession, err = gorm.Open(mysql.Open(ConnectUrl()), &gorm.Config{})
	if err != nil {
		log.Logger.Error("mysql init err")
		return err
	}
	return nil
}

func ConnectUrl() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Config.Mysql.UserName,
		config.Config.Mysql.Password,
		config.Config.Mysql.Url,
		config.Config.Mysql.Port,
		config.Config.Mysql.DbName,
	)
}

func MysqlSession() *gorm.DB {
	return SqlSession
}
