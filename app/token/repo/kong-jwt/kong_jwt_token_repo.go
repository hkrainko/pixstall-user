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

var regSecret = []byte("PTDpYC0h0A1FEch84x3U9G4otA11NzSC")

type RegClaims struct {
	AuthID string `json:"authID"`
	jwt.StandardClaims
}

var apiSecret = []byte("nBcWcVKTRiiUT0iaahFBFskAlugkP5GX")

type APIClaims struct {
	UserID string `json:"userId"`
	jwt.StandardClaims
}

const (
	KongURL = "localhost:8001"
)

type kongJWTTokenRepo struct {
}

func NewKongJWTTokenRepo() token.Repo {
	return &kongJWTTokenRepo{

	}
}

func (k kongJWTTokenRepo) GenerateRegToken(ctx context.Context, authID string) (string, error) {
	// set claims and sign
	claims := RegClaims{
		AuthID: authID,
		StandardClaims: jwt.StandardClaims{
			Audience:  "API",
			ExpiresAt: time.Now().Add(1 * time.Hour).Unix(),
			Id:        "pixstall-reg-jwt",
			IssuedAt:  time.Now().Unix(),
			Issuer:    "pixstall-reg-cred",
			Subject:   authID,
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtToken, err := tokenClaims.SignedString(regSecret)
	if err != nil {
		return "", err
	}
	return jwtToken, nil
}

func (k kongJWTTokenRepo) GenerateAPIToken(ctx context.Context, userID string) (string, error) {
	// set claims and sign
	claims := APIClaims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			Audience:  "API",
			ExpiresAt: time.Now().Add(10 * time.Second).Unix(),
			Id:        "pixstall-api-jwt",
			IssuedAt:  time.Now().Unix(),
			Issuer:    "pixstall-api-cred",
			Subject:   userID,
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtToken, err := tokenClaims.SignedString(apiSecret)
	if err != nil {
		return "", err
	}
	return jwtToken, nil
}
