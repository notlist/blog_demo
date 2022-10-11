package entity

// 数据库表明自定义，默认为model的复数形式，比如这里默认为 users
func (User) TableName() string {
	return "sys_user"
}

type User struct {
	Id         int    `json:"id" gorm:"column:id"`
	UserId     int64  `json:"user_id"  gorm:"column:user_id"`
	Name       string `json:"name" gorm:"column:name"`
	Password   string `json:"password" gorm:"column:password"`
	Email      string `json:"email" gorm:"column:email"`
	DelStatus  int    `json:"del_status" gorm:"column:del_status"`
	UpdateTime int64  `json:"update_time" gorm:"column:update_time"`
	CreateTime int64  `json:"create_time" gorm:"create_time"`
}
