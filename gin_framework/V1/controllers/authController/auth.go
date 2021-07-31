package authController

import "github.com/gin-gonic/gin"

type AuthController interface {
	SingIn(context *gin.Context)
	LogIn(context *gin.Context)
}

type authController struct {
}

func New() {

}
