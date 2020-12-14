package jwt

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"pixstall-user/domain/token"
	"time"
)

type APIClaims struct {
	UserID string `json:"userId"`
	jwt.StandardClaims
}

type RegClaims struct {
	AuthID string `json:"authId"`
	jwt.StandardClaims
}

type RefreshClaims struct {
	UserID string `json:"authId"`
	jwt.StandardClaims
}

var apiJWTSecret = []byte("apiJWTSecret_dummy_key")
var regJWTSecret = []byte("regJWTSecret_dummy_key")
var refreshJWTSecret = []byte("refreshJwtSecret_dummy_key")

type jwtTokenRepo struct {
}

func NewJWTTokenRepo() token.Repo {
	return &jwtTokenRepo{

	}
}

func (j jwtTokenRepo) GenerateAPIToken(ctx context.Context, userID string) (string, error) {
	// set claims and sign
	claims := APIClaims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			Audience:  "API",
			ExpiresAt: time.Now().Add(1 * time.Hour).Unix(),
			Id:        "pixstall-api-jwt",
			IssuedAt:  time.Now().Unix(),
			Issuer:    "pixstall",
			Subject:   userID,
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtToken, err := tokenClaims.SignedString(apiJWTSecret)
	if err != nil {
		return "", err
	}
	return jwtToken, nil
}

func (j jwtTokenRepo) GenerateRegToken(ctx context.Context, authID string) (string, error) {
	// set claims and sign
	claims := RegClaims{
		AuthID: authID,
		StandardClaims: jwt.StandardClaims{
			Audience:  "REG",
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			Id:        "pixstall-reg-jwt",
			IssuedAt:  time.Now().Unix(),
			Issuer:    "pixstall",
			Subject:   authID,
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtToken, err := tokenClaims.SignedString(regJWTSecret)
	if err != nil {
		return "", err
	}
	return jwtToken, nil
}

func (j jwtTokenRepo) GenerateRefreshToken(ctx context.Context, userID string) (string, error) {
	claims := RefreshClaims{
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
	jwtToken, err := tokenClaims.SignedString(refreshJWTSecret)
	if err != nil {
		return "", err
	}
	return jwtToken, nil
}
