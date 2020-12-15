package kong_jwt

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"net/url"
	"pixstall-user/app/token/repo/kong-jwt/model"
	"pixstall-user/domain/token"
)

type RefreshClaims struct {
	UserID string `json:"authId"`
	jwt.StandardClaims
}

var apiSecret = []byte("api_dummy_key")

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
	return "", errors.New("")
}

func retrieveConsumer(consumerUserName string) (int, error) {
	resp, err := http.Get("http://localhost:8001/consumers/" + consumerUserName)
	if err != nil {
		return 0, err
	}
	return resp.StatusCode, nil
}

func addConsumerIfNotExist(consumerUserName string) (string, error) {
	data := url.Values{
		"username":  {consumerUserName},
	}
	resp, err := http.PostForm("http://localhost:8001/consumers", data)
	if err != nil {
		return "", err
	}
	if resp.StatusCode == http.StatusConflict {
		return "", &model.DuplicatedConsumerError{}
	}
	var res map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return "", err
	}
	return res["username"].(string), err
}

func getJWTCredential(consumerUserName string) (*model.JWTCredential, error) {
	resp, err := http.Get("http://localhost:8001/consumers/" + consumerUserName + "/jwt")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("user not found")
	}
	var lJWTCredResp = model.ListJWTCredentialsResponse{}
	err = json.NewDecoder(resp.Body).Decode(&lJWTCredResp)
	if err != nil {
		return nil, err
	}
	if len(lJWTCredResp.Data) < 0 {
		return nil, errors.New("No credential")
	}
	return &lJWTCredResp.Data[0], nil
}

func createJWTCredentialIfNotExist(consumerUserName string) error {

	resp, err := http.Get("http://localhost:8001/consumers/" + consumerUserName + "/jwt")
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return errors.New("user not found")
	}
	var lJWTCredResp = model.ListJWTCredentialsResponse{}
	err = json.NewDecoder(resp.Body).Decode(&lJWTCredResp)
	if err != nil {
		return err
	}
	if len(lJWTCredResp.Data) > 0 {
		return nil
	}

	//create one the return

	return nil
}


