package kong_jwt

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
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

func TestKongJWTTokenRepo_GenerateToken(t *testing.T) {
	dropAll()
	token, err := repo.GenerateToken(ctx, "tempUserID_")
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func dropAll() {

}