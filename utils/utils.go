package utils

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sreehari212000/blog/models"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(passwordString string) (string, error) {
	cost := bcrypt.DefaultCost
	passwordBytes, err := bcrypt.GenerateFromPassword([]byte(passwordString), cost)
	if err != nil {
		return "", err
	}
	hashedPassword := string(passwordBytes)
	return hashedPassword, nil
}

func IsPasswordValid(passwordFromDB string, passwordFromRequest string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(passwordFromDB), []byte(passwordFromRequest))
	return err == nil
}

var key = []byte("mysecretkey")

func CreateJwtToken(claim models.JwtClaim) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":   claim.Email,
		"user_id": claim.User_Id,
	})
	s, err := t.SignedString(key)
	if err != nil {
		return "", err
	}
	return s, nil
}

func VerifyJwtToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return key, nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if err != nil {
		// log.Fatal(err)
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims["user_id"].(string), nil
	} else {
		fmt.Println(err)
		return "", err
	}

}
