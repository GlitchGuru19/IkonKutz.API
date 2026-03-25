package utils

import (
	"time"

	"IkonKutz.API/initializers"
	"github.com/golang-jwt/jwt/v4"
)

// struct for JWT claims
type JWTClaims struct {
	UserID uint   `json:"userId"`
	Role   string `json:"role"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

// fuction to generate JWT token
func GenerateToken(userID uint, role string, email string) (string, error) {
	expiresAt := time.Now().Add(time.Duration(initializers.AppConfig.JWTExpiresInHours) * time.Hour)

	claims := JWTClaims{
		UserID: userID,
		Role:   role,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(initializers.AppConfig.JWTSecret))
}

// function to parse JWT token
func ParseToken(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(initializers.AppConfig.JWTSecret), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JWTClaims)
	if !ok || !token.Valid {
		return nil, jwt.ErrTokenInvalidClaims
	}

	return claims, nil
}
