package authService

import (
	"V1/mocks"
	"V1/models"
	"V1/models/entity"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

func TestSignIn(t *testing.T) {
	type expected struct {
		ok  bool
		err error
	}
	tests := []struct {
		name          string
		input         entity.User
		serviceOutErr error
		expected      expected
	}{
		{"notAge", entity.User{Firstname: "mohammad", Lastname: "dehghanpour"}, nil, expected{false, errors.New("invalid input")}},
		{"notFirstname", entity.User{Age: 20, Lastname: "dehghanpour"}, nil, expected{false, errors.New("invalid input")}},
		{"notLastname", entity.User{Age: 20, Firstname: "mohammad"}, nil, expected{false, errors.New("invalid input")}},
		{"serviceErr", entity.User{Age: 20, Firstname: "mohammad", Lastname: "dehghanpour"}, errors.New("error"), expected{false, errors.New("error")}},
		{"ok", entity.User{Age: 20, Firstname: "mohammad", Lastname: "dehghanpour"}, nil, expected{true, nil}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockRepo := new(mocks.UserCollection)
			mockRepo.On("InsertOne", context.TODO(), test.input).Return(primitive.NilObjectID, test.serviceOutErr)
			service := NewAuthService(mockRepo, nil)

			ok, err := service.SignIn(test.input)

			assert.Equal(t, test.expected.ok, ok)
			if test.expected.err != nil {
				assert.EqualError(t, err, test.expected.err.Error())
			} else {
				assert.Nil(t, err)
			}

		})
	}

}

func TestLogIn(t *testing.T) {
	type repoOut struct {
		user *models.User
		err  error
	}
	type jwtOut struct {
		token string
		err   error
	}
	tests := []struct {
		name          string
		input         entity.UserLogIn
		expectedToken string
		expectedErr   error
		repoOut       repoOut
		jwtOut        jwtOut
	}{
		{"notFirstname", entity.UserLogIn{"", "dehghanpour"}, "", errors.New("error"), repoOut{&models.User{primitive.NewObjectID(), "mohammad", "dehghanpour", 22, nil}, nil}, jwtOut{"token", nil}},
		{"notLastname", entity.UserLogIn{"mohammad", ""}, "", errors.New("error"), repoOut{&models.User{primitive.NewObjectID(), "mohammad", "dehghanpour", 22, nil}, nil}, jwtOut{"token", nil}},
		{"ok", entity.UserLogIn{"mohammad", "dehghanpour"}, "token", nil, repoOut{&models.User{primitive.NewObjectID(), "mohammad", "dehghanpour", 22, nil}, nil}, jwtOut{"token", nil}},
		{"repoErr", entity.UserLogIn{"mohammad", "dehghanpour"}, "", errors.New("error"), repoOut{&models.User{primitive.NewObjectID(), "mohammad", "dehghanpour", 22, nil}, errors.New("error")}, jwtOut{"token", nil}},
		{"jwtErr", entity.UserLogIn{"mohammad", "dehghanpour"}, "", errors.New("error"), repoOut{&models.User{primitive.NewObjectID(), "mohammad", "dehghanpour", 22, nil}, nil}, jwtOut{"token", errors.New("error")}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockRepo := new(mocks.UserCollection)
			mockRepo.On("Aggregate", context.TODO(), mock.Anything).Return(test.repoOut.user, test.repoOut.err)
			mockJWT := new(mocks.JWTService)
			mockJWT.On("Generator", mock.Anything).Return(test.jwtOut.token, test.jwtOut.err)
			auth := NewAuthService(mockRepo, mockJWT)

			token, err := auth.LogIn(test.input)

			if err == nil {
				assert.Nil(t, test.expectedErr)
			} else {
				assert.NotNil(t, test.expectedErr)
			}
			assert.Equal(t, test.expectedToken, token)

		})
	}
}
