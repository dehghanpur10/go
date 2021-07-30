package main

import (
	"V1/routers/videoRouter"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.New()
	server.Use(gin.Recovery(), gin.Logger())
	//server.Use(middleware func)    ==> add custom middleware
	api := server.Group("api/V1")
	videoRouter.SetVideoRouter(api)
	err := server.Run(":8080")
	if err != nil {
		fmt.Println("error occur in run server")
	}

}
