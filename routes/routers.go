package routes

import (
	"github.com/gin-gonic/gin"
	"goadmin/controller"
)

func SetRouter() *gin.Engine {
	r := gin.Default()

	/**
	用户User路由组
	*/
	userGroup := r.Group("user")
	{
		//增加用户User
		userGroup.POST("/add", controller.CreateUser)
		//查看所有的User
		userGroup.GET("/all", controller.GetUserList)
		//修改某个User
		userGroup.POST("/update/", controller.UpdateUser)
		//删除某个User
		userGroup.POST("/delete/:id/", controller.DeleteUserById)
	}

	return r
}
