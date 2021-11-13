package util

import (
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type LoginClaim struct {
	jwt.StandardClaims
	UserId uint64
}

var secret string = "goblogsecret"

//生成JWT Token
func GenJWTToken(userId uint64) (string, error) {
	now := time.Now()
	claim := LoginClaim{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  now.Unix(),
			ExpiresAt: now.Add(time.Hour).Unix(),
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claim).SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return token, nil
}

func VerifyJWTToken(tokenStr string) (*LoginClaim, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &LoginClaim{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		},
	)
	if err != nil {
		return nil, err
	}

	if claim, ok := token.Claims.(*LoginClaim); ok && token.Valid {
		return claim, nil
	}
	return nil, errors.New("invalid token")
}
