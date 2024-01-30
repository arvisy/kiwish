package helper

import (
	"errors"

	"github.com/golang-jwt/jwt/v4"
)

var jwtSecret = []byte("your-secret-key")

func GenerateJWTTokenWithClaims(claims jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", errors.New("failed to sign JWT token")
	}

	return signedToken, nil
}

// func GenerateJWTToken(user *model.User) (string, error) {
// 	claims := jwt.MapClaims{
// 		"id": user.Id,
// 		// "role": user.Role,
// 		"exp": time.Now().Add(time.Hour * 24).Unix(),
// 	}

// 	return GenerateJWTTokenWithClaims(claims)
// }
