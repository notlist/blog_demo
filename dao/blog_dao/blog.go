package blog_dao

import (
	"goadmin/common/log"
	"goadmin/entity"
	"gorm.io/gorm"
)

func (o *BlogImpl) Add(blog *entity.Blog) error {
	res := o.Db.Create(&blog)
	if res == nil || res.Error != nil {
		log.Logger.Errorf("add record err:%+v", res.Error)
		return res.Error
	}
	return nil
}
func (o *BlogImpl) GetAll(cond map[string]interface{}) ([]*entity.Blog, error) {
	list := make([]*entity.Blog, 0)
	res := o.Db.Where(cond).Find(&list)
	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		log.Logger.Errorf("get blogs record err:%+v", res.Error)
		return nil, res.Error
	}
	return list, nil
}
func (o *BlogImpl) GetOne(cond map[string]interface{}) (*entity.Blog, error) {
	info := &entity.Blog{}
	res := o.Db.Where(cond).Find(&info)
	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		log.Logger.Errorf("get blog record err:%+v", res.Error)
		return nil, res.Error
	}
	return info, nil
}
func (o *BlogImpl) Update(cond map[string]interface{}, updateData map[string]interface{}) error {
	res := o.Db.Table("blog").Where(cond).Updates(&updateData)
	if res == nil || res.Error != nil {
		log.Logger.Errorf("Update blog record err:%+v", res.Error)
		return res.Error
	}
	return nil
}
func (o *BlogImpl) Delete(cond map[string]interface{}) error {
	res := o.Db.Where(cond).Delete(&entity.Blog{})
	if res == nil || res.Error != nil {
		log.Logger.Errorf("Update record err:%+v", res.Error)
		return res.Error
	}
	return nil
}
