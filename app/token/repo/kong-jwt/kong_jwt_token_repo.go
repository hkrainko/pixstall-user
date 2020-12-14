package kong_jwt

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"pixstall-user/domain/token"
	"time"
)

type RefreshClaims struct {
	UserID string `json:"authId"`
	jwt.StandardClaims
}

var refreshJWTSecret = []byte("refreshJwtSecret_dummy_key")
const (
	KongURL = "localhost:8001"
)

type kongJWTTokenRepo struct {

}

func NewKongJWTTokenRepo() token.Repo {
	return &kongJWTTokenRepo{

	}
}

func (k kongJWTTokenRepo) GenerateToken(ctx context.Context, userID string) (string, error) {
	panic("implement me")
}

func (k kongJWTTokenRepo) GenerateRefreshToken(ctx context.Context, userID string) (string, error) {
	claims := RefreshClaims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			Audience:  userID,
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			Id:        "pixstall-user-refresh-default",
			IssuedAt:  time.Now().Unix(),
			Issuer:    "pixstall",
			Subject:   userID,
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtToken, err := tokenClaims.SignedString(refreshJWTSecret)
	if err != nil {
		return "", err
	}
	return jwtToken, nil
}