package group

import (
	"github.com/gin-gonic/gin"
	"goadmin/controller"
)

// 用户路由组
func UserGroup(r *gin.Engine) {
	userGroup := r.Group("user/test")
	{
		userGroup.POST("/login", controller.LoginUser) //增加用户User
		userGroup.POST("/all", controller.SignUser)    //查看所有的User
	}

}
