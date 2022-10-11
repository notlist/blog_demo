package controller

import (
	"github.com/gin-gonic/gin"
	"goadmin/common/rsp"
	"goadmin/dto"
	"goadmin/service"
)

func SignUser(c *gin.Context) {
	//定义一个User变量
	var user dto.UserAddReq
	//将调用后端的request请求中的body数据根据json格式解析到User结构变量中
	c.BindJSON(&user)
	//将被转换的user变量传给service层的CreateUser方法，进行User的新建
	err := service.SignUser(&user)
	//判断是否异常，无异常则返回包含200和更新数据的信息
	if err != nil {
		rsp.Error(c, err.Error())
	} else {
		rsp.Success(c, "注册成功", nil)
	}
}

func LoginUser(c *gin.Context) {
	c.Request.Context()
	//定义一个User变量
	var user dto.UserAddReq
	//将调用后端的request请求中的body数据根据json格式解析到User结构变量中
	c.BindJSON(&user)
	//将被转换的user变量传给service层的CreateUser方法，进行User的新建
	err := service.LoginUser(c, &user)
	//判断是否异常，无异常则返回包含200和更新数据的信息
	if err != nil {
		rsp.Error(c, err.Error())
	} else {
		rsp.Success(c, "登录成功", nil)
	}
}
