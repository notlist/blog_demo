package blog_dao

import (
	"goadmin/common/mysql"
	"goadmin/entity"
	"gorm.io/gorm"
)

type BlogDao interface {
	Add(user *entity.Blog) error
	GetAll(cond map[string]interface{}) ([]*entity.Blog, error)
	GetOne(cond map[string]interface{}) (*entity.Blog, error)
	Update(cond map[string]interface{}, updateData map[string]interface{}) error
	Delete(cond map[string]interface{}) error
}

type BlogImpl struct {
	Db *gorm.DB
}

func BlogDaoNew() *BlogImpl {
	return &BlogImpl{
		Db: mysql.MysqlSession(),
	}
}
