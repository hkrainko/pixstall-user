package kong_jwt

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func setup() {
	fmt.Println("Before all tests")
	//ctx = context.Background()
}

func teardown() {
	dropAll()
	fmt.Println("After all tests")
}

func dropAll() {

}