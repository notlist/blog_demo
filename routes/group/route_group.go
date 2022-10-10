package group

import (
	"github.com/gin-gonic/gin"
	"goadmin/controller"
)

// 用户路由组
func UserGroup(r *gin.Engine) {

	group := r.Group("user")
	{
		//增加用户User
		group.POST("/add", controller.CreateUser)
		//查看所有的User
		group.GET("/all", controller.GetUserList)
		//修改某个User
		group.POST("/update/", controller.UpdateUser)
		//删除某个User
		group.POST("/delete/:id/", controller.DeleteUserById)
	}

}
