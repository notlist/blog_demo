package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"goadmin/common/log"
	"goadmin/dao/user_dao"
	"goadmin/entity"
	"goadmin/request"
	// 导入session包
	"github.com/gin-contrib/sessions"

	"time"
)

func SignUser(req *request.UserAddReq) (err error) {
	dao := user_dao.UserDaoNew()

	if req.Name == "" {
		return errors.New("用户名不能为空")
	}
	info, err := dao.GetOne(map[string]interface{}{
		"name": req.Name,
	})
	if err != nil {
		log.Logger.Errorf("get user info err:%+v", err)
		return errors.New("服务器错误")
	}
	if info != nil && info.Name == req.Name {
		return errors.New("该用户名已经存在")
	}

	userInfo := &entity.User{
		UserId:     time.Now().Unix() - 1000000,
		Name:       req.Name,
		Password:   req.Password,
		Email:      req.Email,
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
	}
	return user_dao.UserDaoNew().Add(userInfo)
}

func LoginUser(c *gin.Context, req *request.UserAddReq) (err error) {
	dao := user_dao.UserDaoNew()
	if req.Name == "" || req.Password == "" {
		return errors.New("用户名/密码不能为空")
	}
	info, err := dao.GetOne(map[string]interface{}{
		"name":     req.Name,
		"password": req.Password,
	})
	if err != nil {
		log.Logger.Errorf("get user info err:%+v", err)
		return errors.New("服务器错误")
	}
	if info == nil {
		return errors.New("用户名或密码错误")
	}
	session := sessions.Default(c)
	session.Set(info.UserId, 1)
	return nil
}

///*
//*
//获取user集合
//*/
//func GetAllUser() (userList []*entity.User, err error) {
//
//	return user_dao.SelectList()
//}
//
///*
//*
//根据id删除user
//*/
//func DeleteUserById(id string) (err error) {
//	return user_dao.DeleteById(id)
//}
//
///*
//*
//根据id查询用户User
//*/
//func GetUserById(id int) (user *entity.User, err error) {
//	return user_dao.GetUserById(id)
//}
//
///*
//*
//更新用户信息
//*/
//func UpdateUser(user *entity.User) (err error) {
//	return user_dao.UpdateUser(user)
//}
