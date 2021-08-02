package authController

import (
	"V1/mocks"
	"V1/models/entity"
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSignInController(t *testing.T) {
	const PATH string = "/api/V1/signIn"
	//setup
	gin.SetMode(gin.TestMode)
	type serviceOut struct {
		ok  bool
		err error
	}
	tests := []struct {
		name       string
		user       entity.User
		serviceOut serviceOut
		status     int
		message    interface{}
	}{
		{"ok", entity.User{"mohammad", "dehghanpour", 20}, serviceOut{true, nil}, http.StatusCreated, true},
		{"not ok", entity.User{"mohammad", "dehghanpour", 20}, serviceOut{false, nil}, http.StatusCreated, false},
		{"not firstname", entity.User{Lastname: "dehghanpour", Age: 20}, serviceOut{true, nil}, http.StatusBadRequest, "invalid input data"},
		{"not lastname", entity.User{Firstname: "mohammad", Age: 20}, serviceOut{true, nil}, http.StatusBadRequest, "invalid input data"},
		{"not age", entity.User{Firstname: "mohammad", Lastname: "dehghanpour"}, serviceOut{true, nil}, http.StatusBadRequest, "invalid input data"},
		{"service error", entity.User{Firstname: "mohammad", Lastname: "dehghanpour", Age: 20}, serviceOut{false, errors.New("service error")}, http.StatusBadRequest, "service error"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			service := new(mocks.AuthService)
			service.On("SignIn", mock.AnythingOfType("entity.User")).Return(test.serviceOut.ok, test.serviceOut.err)

			controller := NewAuthController(service)
			// a response recorder for getting written http response
			rr := httptest.NewRecorder()

			router := gin.Default()
			router.POST(PATH, controller.SingIn)

			body, err := json.Marshal(test.user)
			assert.NoError(t, err)

			stream := bytes.NewReader(body)
			request, err := http.NewRequest(http.MethodPost, PATH, stream)
			assert.NoError(t, err)

			router.ServeHTTP(rr, request)

			respBody, err := json.Marshal(map[string]interface{}{
				"message": test.message,
			})
			assert.NoError(t, err)

			assert.Equal(t, test.status, rr.Code)
			assert.Equal(t, respBody, rr.Body.Bytes())

			//service.AssertExpectations(t) // assert that UserService.Get was called

		})
	}

}
func TestLogInController(t *testing.T) {
	const PATH = "/api/V1/login"
	gin.SetMode(gin.TestMode)
	type serviceOut struct {
		token string
		err   error
	}
	tests := []struct {
		name       string
		user       entity.UserLogIn
		serviceOut serviceOut
		status     int
		message    interface{}
	}{
		{"ok", entity.UserLogIn{"mohammad", "dehghanpour"}, serviceOut{"token", nil}, http.StatusOK, "token"},
		{"service error", entity.UserLogIn{"mohammad", "dehghanpour"}, serviceOut{"", errors.New("error")}, http.StatusBadRequest, "error"},
		{"not firstname", entity.UserLogIn{"", "dehghanpour"}, serviceOut{"token", nil}, http.StatusBadRequest, "invalid input data"},
		{"not lastname", entity.UserLogIn{"mohammad", ""}, serviceOut{"token", nil}, http.StatusBadRequest, "invalid input data"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			service := new(mocks.AuthService)
			service.On("LogIn", mock.AnythingOfType("entity.UserLogIn")).Return(test.serviceOut.token, test.serviceOut.err)

			controller := NewAuthController(service)
			// a response recorder for getting written http response
			rr := httptest.NewRecorder()

			router := gin.Default()
			router.POST(PATH, controller.LogIn)

			body, err := json.Marshal(test.user)
			assert.NoError(t, err)

			stream := bytes.NewReader(body)
			request, err := http.NewRequest(http.MethodPost, PATH, stream)
			assert.NoError(t, err)

			router.ServeHTTP(rr, request)

			respBody, err := json.Marshal(map[string]interface{}{
				"message": test.message,
			})
			assert.NoError(t, err)

			assert.Equal(t, test.status, rr.Code)
			assert.Equal(t, respBody, rr.Body.Bytes())

			//service.AssertExpectations(t) // assert that UserService.Get was called

		})
	}
}
