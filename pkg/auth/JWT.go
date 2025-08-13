package auth

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type JWTManager struct {
	secret     string
	expireTime int
}

func NewJWTManager(secret string, expireTime int) *JWTManager {
	return &JWTManager{secret, expireTime}
}

func (j *JWTManager) GenerateToken(userId uint) (string, error) {
	claims := jwt.MapClaims{
		"sub": userId,
		"exp": time.Now().Add(time.Duration(j.expireTime) * time.Hour).Unix(),
		"iat": time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secret))
}

func (j *JWTManager) VerifyToken(tokenStr string) (uint, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(j.secret), nil
	})
	if err != nil || !token.Valid {
		return 0, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok {
		if sub, ok := claims["sub"].(float64); ok {
			return uint(sub), nil
		}
	}
	return 0, errors.New("invalid token")
}
