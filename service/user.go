package service

import (
	"goadmin/dao/user_dao"
	"goadmin/entity"
	"goadmin/request"
)

/*
*
新建User信息
*/
func CreateUser(req *request.UserAddReq) (err error) {
	userInfo := &entity.User{
		Name:     req.Name,
		NickName: req.NickName,
		Avatar:   req.Avatar,
		Password: req.Password,
		Email:    req.Email,
		Mobile:   req.Mobile,
	}
	return user_dao.UserDaoNew().Add(userInfo)
}

/*
*
获取user集合
*/
func GetAllUser() (userList []*entity.User, err error) {

	return user_dao.SelectList()
}

/*
*
根据id删除user
*/
func DeleteUserById(id string) (err error) {
	return user_dao.DeleteById(id)
}

/*
*
根据id查询用户User
*/
func GetUserById(id int) (user *entity.User, err error) {
	return user_dao.GetUserById(id)
}

/*
*
更新用户信息
*/
func UpdateUser(user *entity.User) (err error) {
	return user_dao.UpdateUser(user)
}
