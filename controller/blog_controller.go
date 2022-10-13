package controller

import (
	"errors"
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

	resp, err := service.BlogList(req.UserId, &req)
	if err != nil {
		rsp.Error(c, err.Error())
	} else {
		rsp.Success(c, "success", resp)
	}
}

func CreateBlog(c *gin.Context) {
	var req dto.BlogCreateReq

	c.BindJSON(&req)

	userId := util.GetCurrentUser(c)
	if userId == "" {
		rsp.Error(c, errors.New("请先登录").Error())
		return
	}
	uid, _ := strconv.ParseInt(userId, 10, 64)

	err := service.CreateBlog(uid, &req)
	if err != nil {
		rsp.Error(c, err.Error())
	} else {
		rsp.Success(c, "success", nil)
	}

}

func EditBlog(c *gin.Context) {
	var req dto.BlogEditReq

	c.BindJSON(&req)

	userId := util.GetCurrentUser(c)
	if userId == "" {
		rsp.Error(c, errors.New("请先登录").Error())
		return
	}
	uid, _ := strconv.ParseInt(userId, 10, 64)

	err := service.EditBlog(uid, &req)
	if err != nil {
		rsp.Error(c, err.Error())
	} else {
		rsp.Success(c, "success", nil)
	}

}

func DeleteBlog(c *gin.Context) {

}

func BLogDetail(c *gin.Context) {

}
