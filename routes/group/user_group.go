package group

import (
	"github.com/gin-gonic/gin"
	"goadmin/controller"
)

// 用户路由组
func UserGroup(r *gin.Engine) {
	userGroup := r.Group("user")
	{
		userGroup.POST("/login", controller.LoginUser) //登录
		userGroup.POST("/sign", controller.SignUser)   //注册
	}

}
