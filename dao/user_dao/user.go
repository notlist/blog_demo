package user_dao

import (
	"goadmin/common/log"
	"goadmin/common/mysql"
	"goadmin/entity"
)

func SelectList() (userList []*entity.User, err error) {
	db := mysql.MysqlSession()
	res := db.Find(&userList)
	if res == nil || res.Error != nil {
		return nil, err
	}
	return
}

func DeleteById(id string) (err error) {
	db := mysql.MysqlSession()
	res := db.Where("id=?", id).Delete(&entity.User{})
	if res == nil || res.Error != nil {
		return err
	}
	return err
}

func GetUserById(id int) (user *entity.User, err error) {
	db := mysql.MysqlSession()
	user = &entity.User{}
	res := db.Where("id=?", id).First(user)
	if res == nil || res.Error != nil {
		return nil, err
	}
	return
}

func UpdateUser(user *entity.User) (err error) {
	db := mysql.MysqlSession()
	err = db.Save(user).Error
	return
}

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
		log.Logger.Errorf("add record err:%+v", res.Error)
		return user, res.Error
	}
	return nil, nil
}
func (o *UserImpl) Update(cond map[string]interface{}, updateData entity.User) error {
	res := o.db.Model(entity.User{}).Where(cond).Update(&updateData)
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
