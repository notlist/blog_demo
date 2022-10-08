package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"goadmin/common/config"
	"goadmin/routes"
)

func main() {
	//连接数据库
	err := config.InitMySql()
	if err != nil {
		panic(err)
	}
	//程序退出关闭数据库连接
	defer config.Close()
	//绑定模型
	//config.SqlSession.AutoMigrate(&entity.User{})
	//注册路由
	r := routes.SetRouter()
	//启动端口为8081的项目
	r.Run(":8081")
}
