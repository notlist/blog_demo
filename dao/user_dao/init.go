package user_dao

import (
	"github.com/jinzhu/gorm"
	"goadmin/common/mysql"
	"goadmin/entity"
)

type UserDao interface {
	Add(user *entity.User) error
	GetAll(cond map[string]interface{}) ([]*entity.User, error)
	GetOne(cond map[string]interface{}) (*entity.User, error)
	Update(uid int64, cond map[string]interface{}, updateData entity.User) error
	Delete(uid int64, cond map[string]interface{}) error
}

type UserImpl struct {
	db *gorm.DB
}

func UserDaoNew() *UserImpl {
	return &UserImpl{
		db: mysql.MysqlSession(),
	}
}
