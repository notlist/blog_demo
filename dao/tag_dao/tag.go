package tag_dao

import (
	"goadmin/common/log"
	"goadmin/entity"
	"gorm.io/gorm"
)

func (o *TagImpl) Add(data *entity.Tag) error {
	res := o.db.Create(data)
	if res == nil || res.Error != nil {
		log.Logger.Errorf("add record err:%+v", res.Error)
		return res.Error
	}
	return nil
}

func (o *TagImpl) BatchAdd(data []*entity.Tag) error {
	res := o.db.CreateInBatches(data, 100)
	if res == nil || res.Error != nil {
		log.Logger.Errorf("batch add record err:%+v", res.Error)
		return res.Error
	}
	return nil
}

func (o *TagImpl) GetAll(cond map[string]interface{}) ([]*entity.Tag, error) {
	list := make([]*entity.Tag, 0)
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

func (o *TagImpl) GetOne(cond map[string]interface{}) (*entity.Tag, error) {
	info := &entity.Tag{}
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

func (o *TagImpl) Update(cond map[string]interface{}, data map[string]interface{}) error {
	res := o.db.Table("tag").Where(cond).Updates(data)
	if res == nil || res.Error != nil {
		log.Logger.Errorf("Update record err:%+v", res.Error)
		return res.Error
	}
	return nil
}

func (o *TagImpl) Delete(cond map[string]interface{}) error {
	res := o.db.Where(cond).Delete(&entity.Tag{})
	if res == nil || res.Error != nil {
		log.Logger.Errorf("Update record err:%+v", res.Error)
		return res.Error
	}
	return nil
}
