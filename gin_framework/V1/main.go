package main

import (
	_ "V1/docs"
	"V1/routers/videoRouter"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title test gin framework
// @version 1.0
// @description This is a restapi for video management
// @contact.name mohammad dehghanpour
// @contact.email m.dehghanpour
// @host http://localhost:8080
// @BasePath /api/V1
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	server := gin.New()
	server.Use(gin.Recovery(), gin.Logger())
	//server.Use(middleware func)    ==> add custom middleware
	api := server.Group("api/V1")
	//set router
	videoRouter.SetVideoRouter(api)
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//run server
	err := server.Run(":8080")
	if err != nil {
		fmt.Println("error occur in run server")
	}

}
