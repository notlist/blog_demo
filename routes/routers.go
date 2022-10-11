package routes

import (
	"github.com/gin-gonic/gin"
	"goadmin/controller"
	"goadmin/middleware"
	"goadmin/routes/group"
)

func SetRouter() *gin.Engine {
	//r := gin.Default()
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.GinBodyLogMiddleware())

	group.UserGroup(r)

	/**
	用户User路由组
	*/
	userGroup := r.Group("user")
	{
		//增加用户User
		userGroup.POST("/login", controller.LoginUser)
		//查看所有的User
		userGroup.POST("/sign", controller.SignUser)

	}

	return r
}
