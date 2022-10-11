package controller

import (
	"github.com/gin-gonic/gin"
	"goadmin/common/rsp"
	"goadmin/common/util"
	"goadmin/dto"
	"goadmin/service"
	"strconv"
)

func BlogList(c *gin.Context) {
	//定义一个User变量
	var req dto.BlogListReq
	//将调用后端的request请求中的body数据根据json格式解析到User结构变量中
	c.BindJSON(&req)

	userId := util.GetCurrentUser(c)
	uid, _ := strconv.ParseInt(userId, 10, 64)

	resp, err := service.BlogList(uid, &req)
	if err != nil {
		rsp.Error(c, err.Error())
	} else {
		rsp.Success(c, "success", resp)
	}
}
