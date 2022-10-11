package util

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func GetCurrentUser(c *gin.Context) string {
	session := sessions.Default(c)
	info := session.Get("userId")
	userInfo := info.(string) // 类型转换一下
	return userInfo
}

func SetCurrentUser(c *gin.Context, userId string) {
	session := sessions.Default(c)
	session.Set("userId", userId)
	// 一定要Save否则不生效，若未使用gob注册User结构体，调用Save时会返回一个Error
	session.Save()
}
