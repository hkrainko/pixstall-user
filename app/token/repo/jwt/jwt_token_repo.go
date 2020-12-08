package jwt

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"pixstall-user/domain/token"
	"time"
)

type Claims struct {
	UserID string `json:"userId"`
	jwt.StandardClaims
}

var jwtSecret = []byte("dummy_key")
var refreshJwtSecret = []byte("refreshJwtSecret_dummy_key")

type jwtTokenRepo struct {
}

func NewJWTTokenRepo() token.Repo {
	return &jwtTokenRepo{

	}
}

func (j jwtTokenRepo) GenerateAuthToken(ctx context.Context, userID string) (string, error) {
	// set claims and sign
	claims := Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			Audience:  userID,
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			Id:        "pixstall-user-jwt",
			IssuedAt:  time.Now().Unix(),
			Issuer:    "pixstall",
			Subject:   userID,
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtToken, err := tokenClaims.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return jwtToken, nil
}

func (j jwtTokenRepo) GenerateRefreshToken(ctx context.Context, userID string) (string, error) {
	claims := Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			Audience:  userID,
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			Id:        "pixstall-user-refresh-jwt",
			IssuedAt:  time.Now().Unix(),
			Issuer:    "pixstall",
			Subject:   userID,
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtToken, err := tokenClaims.SignedString(refreshJwtSecret)
	if err != nil {
		return "", err
	}
	return jwtToken, nil
}
