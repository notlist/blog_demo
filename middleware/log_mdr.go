package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"goadmin/common/log"
)

// 获取返回体的中间件
func GinBodyLogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		blw := &log.BodyLogWriter{Body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw
		c.Next()
		responseStr := blw.Body.String()

		//请求方式
		reqMethod := c.Request.Method
		//请求路由
		reqUri := c.Request.RequestURI
		// 状态码
		statusCode := c.Writer.Status()
		//请求ip
		clientIP := c.ClientIP()
		//请求参数
		reqParams := c.Request.Body
		//请求ua
		reqUa := c.Request.UserAgent()
		var resultBody logrus.Fields
		resultBody = make(map[string]interface{})
		resultBody["response"] = responseStr
		resultBody["requestUri"] = reqUri
		resultBody["clientIp"] = clientIP
		resultBody["reqParams"] = reqParams
		resultBody["userAgent"] = reqUa
		resultBody["requestMethod"] = reqMethod
		resultBody["statusCode"] = statusCode
		log.LogErrorInfoToFile(resultBody)
		log.SetUid()
	}
}
