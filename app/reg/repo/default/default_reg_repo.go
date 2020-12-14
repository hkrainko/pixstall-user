package _default

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"pixstall-user/domain/reg"
	"time"
)

type RegClaims struct {
	AuthID string `json:"authId"`
	jwt.StandardClaims
}

var regJWTSecret = []byte("regJWTSecret_dummy_key")

type defaultRegRepo struct {
}

func NewDefaultRegRepo() reg.Repo {
	return &defaultRegRepo{

	}
}

func (d defaultRegRepo) GenerateRegToken(ctx context.Context, authID string) (string, error) {
	// set claims and sign
	claims := RegClaims{
		AuthID: authID,
		StandardClaims: jwt.StandardClaims{
			Audience:  "REG",
			ExpiresAt: time.Now().Add(1 * time.Hour).Unix(),
			Id:        "pixstall-reg-default",
			IssuedAt:  time.Now().Unix(),
			Issuer:    authID,
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtToken, err := tokenClaims.SignedString(regJWTSecret)
	if err != nil {
		return "", err
	}
	return jwtToken, nil
}

func (d defaultRegRepo) ValidateRegToken(ctx context.Context, token string) (string, error) {
	panic("implement me")
}