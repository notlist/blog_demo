package tag_dao

import (
	"goadmin/common/mysql"
	"goadmin/entity"
	"gorm.io/gorm"
)

type TagDao interface {
	Add(data *entity.Tag) error
	BatchAdd(data []*entity.Tag) error
	GetAll(cond map[string]interface{}) ([]*entity.Tag, error)
	GetOne(cond map[string]interface{}) (*entity.Tag, error)
	Update(cond map[string]interface{}, updateData map[string]interface{}) error
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
