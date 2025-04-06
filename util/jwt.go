package util

import (
	"fmt"
	"time"

	"github.com/GabriellaAmah/go-url-shortner/config"
	"github.com/golang-jwt/jwt/v5"
)


type JWTClaims struct {
	jwt.RegisteredClaims
	UserId string `json:"userId"`  
	Email string `json:"email"`  
	Username string `json:"username"`  
}

func CreateJwtToken(data map[string]string) (string, error) {

	claims  := 	&JWTClaims{
		UserId: data["id"],
		Email: data["email"],
		Username: data["username"],
		RegisteredClaims: jwt.RegisteredClaims{
			// Set the exp and sub claims. sub is usually the userID
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			Subject:   "Bearer token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, error := token.SignedString([]byte(config.EnvData.JWT_SECRET_KEY))
	if error != nil {
		return "", error
	}

	return signedToken, nil
}

func DecodeAuthJwt(token string) (JWTClaims, MakeError)   {
	claims := &JWTClaims{}
	parsedToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v",       token.Header["alg"])
			}
		return []byte(config.EnvData.JWT_SECRET_KEY), nil
	})
	
	if err != nil || !parsedToken.Valid {
		return *claims, MakeError{
			IsError: true,
			Message: "Invalid Bearer Token",
			Code:    401,
		}
	} 

	return *claims, MakeError{}
}
