package entity

// 数据库表明自定义，默认为model的复数形式，比如这里默认为 users
func (Tag) TableName() string {
	return "blog"
}

type Tag struct {
	Id         int    `json:"id" gorm:"column:id"`
	Userid     int64  `json:"user_id"  gorm:"column:user_id"`  //用户id
	TagName    string `json:"tag_name" gorm:"column:tag_name"` //标签名称
	BlogId     int64  `json:"blog_id"  gorm:"column:blog_id"`  //博客id
	UpdateTime int64  `json:"update_time" gorm:"column:update_time"`
	CreateTime int64  `json:"create_time" gorm:"create_time"`
}
