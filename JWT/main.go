package main

import (
	"errors"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

// createToken is method that create json web token(JWT)
func createToken(username, password string) (string, error) {
	//create map claims
	claims := jwt.MapClaims{}
	//set data
	claims["username"] = username
	claims["password"] = password
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	//creat token without signed
	T := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	//signed token
	token, err := T.SignedString([]byte("super secret text"))
	if err != nil {
		return "", errors.New("error in signed")
	}
	return token, nil
}
func extractToken(token string) (bool, map[string]interface{}, error) {
	keyFunction := func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("super secret text"), nil
	}
	T, err := jwt.Parse(token, keyFunction)
	if err != nil {
		return false, make(map[string]interface{}), err
	}
	claims, ok := T.Claims.(jwt.MapClaims)
	return ok, claims, nil

}
func main() {

	token, err := createToken("mohammad", "1739")
	fmt.Println(token)
	fmt.Println(err)

	ok, cli, err := extractToken(token)
	fmt.Println(ok)
	fmt.Println(cli)
	fmt.Println(err)

}
