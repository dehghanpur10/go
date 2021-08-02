package JWTService

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

type JWTService interface {
	Generator(map[string]interface{}) (string, error)
	Validator(token string) (bool, map[string]interface{}, error)
	SetDependencyTest()
}

type jwtService struct {
	key            string
	dependencyTest error
}

func NewJWTService(key string) JWTService {
	return &jwtService{
		key:            key,
		dependencyTest: nil,
	}
}

//SetDependencyTest is method to set object for testing
func (service *jwtService) SetDependencyTest() {
	service.dependencyTest = errors.New("dependency error")
}

//Generator is method for generate token with JWT
func (service *jwtService) Generator(data map[string]interface{}) (token string, err error) {

	claims := jwt.MapClaims{}
	for key, value := range data {
		claims[key] = value
	}

	T := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	token, err = T.SignedString([]byte(service.key))

	//for testing
	if service.dependencyTest != nil {
		err = errors.New("dependency Err")
	}

	if err != nil {
		return "", err
	}
	return token, nil
}
func (service *jwtService) Validator(token string) (bool, map[string]interface{}, error) {

	keyFunction := func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(service.key), nil
	}
	T, err := jwt.Parse(token, keyFunction)

	//for testing
	if service.dependencyTest != nil {
		err = errors.New("dependency Err")
	}

	if err != nil {
		return false, make(map[string]interface{}), err
	}
	claims, ok := T.Claims.(jwt.MapClaims)
	return ok, claims, nil
}
