package service

import (
	"goadmin/dao"
	"goadmin/entity"
)

/*
*
新建User信息
*/
func CreateUser(user *entity.User) (err error) {
	return dao.Create(user)
}

/*
*
获取user集合
*/
func GetAllUser() (userList []*entity.User, err error) {

	return dao.SelectList()
}

/*
*
根据id删除user
*/
func DeleteUserById(id string) (err error) {
	return dao.DeleteById(id)
}

/*
*
根据id查询用户User
*/
func GetUserById(id int) (user *entity.User, err error) {
	return dao.GetUserById(id)
}

/*
*
更新用户信息
*/
func UpdateUser(user *entity.User) (err error) {
	return dao.UpdateUser(user)
}
