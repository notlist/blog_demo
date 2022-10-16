package group

import (
	"github.com/gin-gonic/gin"
	"goadmin/controller"
)

// 用户路由组
func BlogGroup(r *gin.Engine) {
	blogGroup := r.Group("blog")
	{
		blogGroup.POST("/list", controller.BlogList)     //列表
		blogGroup.POST("/detail", controller.BLogDetail) //详情
		blogGroup.POST("/add", controller.CreateBlog)    //新增
		blogGroup.POST("/edit", controller.EditBlog)     //编辑
		blogGroup.POST("/delete", controller.DeleteBlog) //删除

	}
}
