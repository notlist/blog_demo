package routes

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"goadmin/middleware"
	"goadmin/routes/group"
)

func SetRouter() *gin.Engine {
	//r := gin.Default()
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.GinBodyLogMiddleware())
	// 创建基于cookie的存储引擎，secret11111 参数是用于加密的密钥
	store := cookie.NewStore([]byte("secret11111"))
	// 设置session中间件，参数mysession，指的是session的名字，也是cookie的名字
	// store是前面创建的存储引擎，我们可以替换成其他存储引擎
	r.Use(sessions.Sessions("mysession", store))

	group.UserGroup(r)
	group.BlogGroup(r)

	return r
}
