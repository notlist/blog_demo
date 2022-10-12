package user_dao

import (
	"goadmin/common/mysql"
	"goadmin/entity"
	"gorm.io/gorm"
)

type UserDao interface {
	Add(user *entity.User) error
	GetAll(cond map[string]interface{}) ([]*entity.User, error)
	GetOne(cond map[string]interface{}) (*entity.User, error)
	Update(cond map[string]interface{}, updateData entity.User) error
	Delete(cond map[string]interface{}) error
}

type UserImpl struct {
	db *gorm.DB
}

func UserDaoNew() *UserImpl {
	return &UserImpl{
		db: mysql.MysqlSession(),
	}
}
