package kong_jwt

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"pixstall-user/domain/token"
	"testing"
)

var repo token.Repo
var ctx context.Context

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func setup() {
	fmt.Println("Before all tests")
	ctx = context.Background()
	repo = NewKongJWTTokenRepo()
}

func teardown() {
	dropAll()
	fmt.Println("After all tests")
}

func TestKongJWTTokenRepo_GenerateAPIToken(t *testing.T) {
	dropAll()
	token, err := repo.GenerateAPIToken(ctx, "temp_UserID")
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	log.Println(token)
}

func TestKongJWTTokenRepo_GenerateRegToken(t *testing.T) {
	dropAll()
	token, err := repo.GenerateRegToken(ctx, "temp_AuthID")
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	log.Println(token)
}

func dropAll() {

}