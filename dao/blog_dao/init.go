package blog_dao

import (
	"github.com/jinzhu/gorm"
	"goadmin/common/mysql"
	"goadmin/entity"
)

type BlogDao interface {
	Add(user *entity.Blog) error
	GetAll(cond map[string]interface{}) ([]*entity.Blog, error)
	GetOne(cond map[string]interface{}) (*entity.Blog, error)
	Update(cond map[string]interface{}, updateData entity.Blog) error
	Delete(cond map[string]interface{}) error
}

type BlogImpl struct {
	db *gorm.DB
}

func BlogDaoNew() *BlogImpl {
	return &BlogImpl{
		db: mysql.MysqlSession(),
	}
}
