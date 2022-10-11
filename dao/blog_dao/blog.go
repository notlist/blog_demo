package blog_dao

import (
	"github.com/jinzhu/gorm"
	"goadmin/common/log"
	"goadmin/entity"
)

func (o *BlogImpl) Add(user *entity.Blog) error {
	res := o.db.Create(user)
	if res == nil || res.Error != nil {
		log.Logger.Errorf("add record err:%+v", res.Error)
		return res.Error
	}
	return nil
}
func (o *BlogImpl) GetAll(cond map[string]interface{}) ([]*entity.Blog, error) {
	list := make([]*entity.Blog, 0)
	res := o.db.Where(cond).Find(&list)
	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		log.Logger.Errorf("get record err:%+v", res.Error)
		return nil, res.Error
	}
	return list, nil
}
func (o *BlogImpl) GetOne(cond map[string]interface{}) (*entity.Blog, error) {
	info := &entity.Blog{}
	res := o.db.Where(cond).Find(&info)
	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		log.Logger.Errorf("get record err:%+v", res.Error)
		return nil, res.Error
	}
	return info, nil
}
func (o *BlogImpl) Update(cond map[string]interface{}, updateData entity.Blog) error {
	res := o.db.Table("blog").Where(cond).Update(&updateData)
	if res == nil || res.Error != nil {
		log.Logger.Errorf("Update record err:%+v", res.Error)
		return res.Error
	}
	return nil
}
func (o *BlogImpl) Delete(cond map[string]interface{}) error {
	res := o.db.Where(cond).Delete(&entity.Blog{})
	if res == nil || res.Error != nil {
		log.Logger.Errorf("Update record err:%+v", res.Error)
		return res.Error
	}
	return nil
}
