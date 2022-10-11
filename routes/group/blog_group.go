package group

import (
	"github.com/gin-gonic/gin"
	"goadmin/controller"
)

// 用户路由组
func BlogGroup(r *gin.Engine) {
	blogGroup := r.Group("blog")
	{
		blogGroup.GET("/list", controller.BlogList)     //列表
		blogGroup.GET("/detail", controller.SignUser)   //详情
		blogGroup.POST("/add", controller.LoginUser)    //新增
		blogGroup.POST("/edit", controller.LoginUser)   //编辑
		blogGroup.POST("/delete", controller.LoginUser) //删除

	}
}
