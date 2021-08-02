package videoRouter

import "github.com/gin-gonic/gin"

func SetVideoRouter(group *gin.RouterGroup) {
	group.GET("video/:videoId")
	group.POST("video")
}
