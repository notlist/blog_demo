package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"goadmin/common/log"
	"goadmin/common/mysql"
	"goadmin/routes"
)

func main() {
	//初始化数据库
	err := mysql.InitMySql()
	if err != nil {
		panic(err)
	}
	//初始化log
	log.Init()
	ssv := "宋思薇"
	log.Logger.Errorf("error is:%s", ssv)
	//程序退出关闭数据库连接
	defer mysql.Close()
	//绑定模型
	//mysql.SqlSession.AutoMigrate(&entity.User{})
	//注册路由
	r := routes.SetRouter()
	//启动端口为8081的项目
	r.Run(":8081")
}
