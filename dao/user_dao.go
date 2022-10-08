package dao

import (
	"goadmin/common/config"
	"goadmin/entity"
)

func Create(user *entity.User) (err error) {
	db := config.MysqlSession()
	res := db.Create(user)
	if res == nil || res.Error != nil {
		return err
	}
	return nil
}

func SelectList() (userList []*entity.User, err error) {
	db := config.MysqlSession()
	res := db.Find(&userList)
	if res == nil || res.Error != nil {
		return nil, err
	}
	return
}

func DeleteById(id string) (err error) {
	db := config.MysqlSession()
	res := db.Where("id=?", id).Delete(&entity.User{})
	if res == nil || res.Error != nil {
		return err
	}
	return err
}

func GetUserById(id int) (user *entity.User, err error) {
	db := config.MysqlSession()
	user = &entity.User{}
	res := db.Where("id=?", id).First(user)
	if res == nil || res.Error != nil {
		return nil, err
	}
	return
}

func UpdateUser(user *entity.User) (err error) {
	db := config.MysqlSession()
	err = db.Save(user).Error
	return
}
