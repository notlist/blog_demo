package dto

type BlogListResp struct {
	BlogId     int64  `json:"blog_id"  gorm:"column:blog_id"` //博客id
	Title      string `json:"title" gorm:"column:title"`      //博客题目
	Content    string `json:"content" gorm:"column:content"`  //博客内容
	UpdateTime int64  `json:"update_time" gorm:"column:update_time"`
	CreateTime int64  `json:"create_time" gorm:"create_time"`
}
