package tag_dao

import (
	"github.com/jinzhu/gorm"
	"goadmin/common/mysql"
	"goadmin/entity"
)

type TagDao interface {
	Add(user *entity.Tag) error
	GetAll(cond map[string]interface{}) ([]*entity.Tag, error)
	GetOne(cond map[string]interface{}) (*entity.Tag, error)
	Update(cond map[string]interface{}, updateData entity.Tag) error
	Delete(cond map[string]interface{}) error
}

type TagImpl struct {
	db *gorm.DB
}

func TagDaoNew() *TagImpl {
	return &TagImpl{
		db: mysql.MysqlSession(),
	}
}
