package authRouter

import "github.com/gin-gonic/gin"

func SetAuthRouter(group *gin.RouterGroup) {
	group.POST("signIn")
	group.POST("login")
}
