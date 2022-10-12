package user_dao

import (
	"github.com/jinzhu/gorm"
	"goadmin/common/log"
	"goadmin/entity"
)

func (o *UserImpl) Add(user *entity.User) error {
	res := o.db.Create(user)
	if res == nil || res.Error != nil {
		log.Logger.Errorf("add record err:%+v", res.Error)
		return res.Error
	}
	return nil
}
func (o *UserImpl) GetAll(cond map[string]interface{}) ([]*entity.User, error) {
	userList := make([]*entity.User, 0)
	res := o.db.Where(cond).Find(&userList)
	if res == nil || res.Error != nil {
		log.Logger.Errorf("add record err:%+v", res.Error)
		return userList, res.Error
	}
	return nil, nil
}
func (o *UserImpl) GetOne(cond map[string]interface{}) (*entity.User, error) {
	user := &entity.User{}
	res := o.db.Where(cond).Find(&user)
	if res == nil || res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		log.Logger.Errorf("get record err:%+v", res.Error)
		return nil, res.Error
	}
	return user, nil
}
func (o *UserImpl) Update(cond map[string]interface{}, updateData entity.User) error {
	res := o.db.Table("sys_user").Where(cond).Updates(&updateData)
	if res == nil || res.Error != nil {
		log.Logger.Errorf("Update record err:%+v", res.Error)
		return res.Error
	}
	return nil
}
func (o *UserImpl) Delete(cond map[string]interface{}) error {
	res := o.db.Where(cond).Delete(&entity.User{})
	if res == nil || res.Error != nil {
		log.Logger.Errorf("Update record err:%+v", res.Error)
		return res.Error
	}
	return nil
}
