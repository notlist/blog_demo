package entity

// 数据库表明自定义，默认为model的复数形式，比如这里默认为 users
func (Blog) TableName() string {
	return "blog"
}

type Blog struct {
	Id         int    `json:"id" gorm:"column:id"`
	BlogId     int64  `json:"blog_id"  gorm:"column:blog_id"` //博客id
	Userid     int64  `json:"user_id"  gorm:"column:user_id"` //用户id
	Title      string `json:"title" gorm:"column:title"`      //博客题目
	Content    string `json:"content" gorm:"column:content"`  //博客内容
	UpdateTime int64  `json:"update_time" gorm:"column:update_time"`
	CreateTime int64  `json:"create_time" gorm:"create_time"`
}
