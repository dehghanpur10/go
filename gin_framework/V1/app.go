package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.New()
	server.Use(gin.Recovery(), gin.Logger())
	//server.Use(middleware func)    ==> add custom middleware

}
