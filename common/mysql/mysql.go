package mysql

import (
	"fmt"
	"goadmin/common/log"
	"gopkg.in/yaml.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io/ioutil"
)

const DRIVER = "mysql"

var SqlSession *gorm.DB

type Conf struct {
	Mysql struct {
		Url      string `yaml:"url"`
		UserName string `yaml:"userName"`
		Password string `yaml:"password"`
		DbName   string `yaml:"dbname"`
		Port     string `yaml:"post"`
	}
}

func (c *Conf) getConf() *Conf {
	yamlFile, err := ioutil.ReadFile("application.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Println(err.Error())
	}
	return c
}

func InitMySql() (err error) {
	var c Conf
	SqlSession, err = gorm.Open(mysql.Open(c.ConnectUrl()), &gorm.Config{})
	if err != nil {
		log.Logger.Error("mysql init err")
		return err
	}
	return nil
}

func (c Conf) ConnectUrl() string {
	conf := c.getConf()
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.Mysql.UserName,
		conf.Mysql.Password,
		conf.Mysql.Url,
		conf.Mysql.Port,
		conf.Mysql.DbName,
	)
}

func MysqlSession() *gorm.DB {
	return SqlSession
}
