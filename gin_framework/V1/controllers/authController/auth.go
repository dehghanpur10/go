package authController

import (
	"V1/models/entity"
	"V1/services/authService"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthController interface {
	SingIn(context *gin.Context)
	LogIn(context *gin.Context)
}

type authController struct {
	authService authService.AuthService
}

func NewAuthController(authService authService.AuthService) AuthController {
	return &authController{
		authService: authService,
	}
}

// SingIn
// @Summary signIn endpoint
// @Description signIn endpoint
// @Tags Authentication
// @Accept json
// @Produce json
// @Param userInfo body entity.User true "user info for signIN"
// @Success 201 {object} entity.User
// @Failure 400 {object} entity.User
// @Router /signIn [post]
func (a *authController) SingIn(context *gin.Context) {
	var user entity.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, map[string]interface{}{"message": "invalid input data"})
		return
	}
	ok, err := a.authService.SignIn(user)
	if err != nil {
		context.JSON(http.StatusBadRequest, map[string]interface{}{"message": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, map[string]interface{}{"message": ok})
}

// LogIn
// @Summary logIn endpoint
// @Description logIn endpoint
// @Tags Authentication
// @Accept json
// @Produce json
// @Param userInfo body entity.UserLogIn true "user info for signIN"
// @Success 201 {object} entity.User
// @Failure 400 {object} entity.User
// @Router /login [post]
func (a *authController) LogIn(context *gin.Context) {
	var user entity.UserLogIn
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, map[string]interface{}{"message": "invalid input data"})
		return
	}
	token, err := a.authService.LogIn(user)
	if err != nil {
		context.JSON(http.StatusBadRequest, map[string]interface{}{"message": err.Error()})
		return
	}
	context.JSON(http.StatusOK, map[string]interface{}{"message": token})
}
