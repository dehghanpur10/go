package authRouter

import (
	authController2 "V1/controllers/authController"
	"V1/repository"
	"V1/repository/userCollection"
	"V1/services/JWTService"
	authService2 "V1/services/authService"
	"github.com/gin-gonic/gin"
)

func SetAuthRouter(group *gin.RouterGroup) {
	var (
		jwtService     = JWTService.NewJWTService("super_secret_ky")
		userCOl        = userCollection.NewUserDatabase(repository.Database)
		authService    = authService2.NewAuthService(userCOl, jwtService)
		authController = authController2.NewAuthController(authService)
	)

	group.POST("signIn", authController.SingIn)

	group.POST("login", authController.LogIn)
}
