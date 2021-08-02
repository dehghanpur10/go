package authService

import (
	"V1/models/entity"
	"V1/repository/userCollection"
	"V1/services/JWTService"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthService interface {
	SignIn(user entity.User) (bool, error)
	LogIn(user entity.UserLogIn) (string, error)
}
type authService struct {
	userCollection userCollection.UserCollection
	JWT            JWTService.JWTService
}

func NewAuthService(userCollection userCollection.UserCollection, JWT JWTService.JWTService) AuthService {
	return &authService{
		userCollection: userCollection,
		JWT:            JWT,
	}
}

//SignIn is method for signin user
func (service *authService) SignIn(user entity.User) (bool, error) {
	if user.Firstname == "" || user.Lastname == "" || user.Age == 0 {
		return false, errors.New("invalid input")
	}
	_, err := service.userCollection.InsertOne(context.TODO(), user)
	if err != nil {
		return false, err
	}
	return true, nil
}

//LogIn is method for login user
func (service *authService) LogIn(user entity.UserLogIn) (string, error) {
	if user.Firstname == "" || user.Lastname == "" {
		return "", errors.New("invalid input")
	}
	// fetch user from database
	filter := bson.D{{"$match", bson.D{{"firstname", user.Firstname}, {"lastname", user.Lastname}}}}
	pipeline := mongo.Pipeline{filter}
	aggregate, err := service.userCollection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		return "", err
	}
	//generate token
	data := make(map[string]interface{})
	data["_id"] = aggregate.ID
	data["firstname"] = aggregate.Firstname
	data["lastname"] = aggregate.Lastname
	generator, err := service.JWT.Generator(data)
	if err != nil {
		return "", err
	}
	return generator, nil
}
