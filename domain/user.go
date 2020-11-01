package domain

import (
	"pixstall_server/app/domain/enums/gender"
	"pixstall_server/app/domain/enums/userstate"
	"time"
)

type User struct {
	ID           string              `json:"id"`
	Token        string              `json:"token"`
	Name         string              `json:"name"`
	AccountID    string              `json:"accountId"`
	Age          int                 `json:"age"`
	Gender       gender.Gender       `json:"gender"`
	State        userstate.UserState `json:"state"`
	LastUpdateDT time.Time           `json:"lastUpdateDt"`
}

type UserRepository interface {
	FindByID(userID string) (*User, error)
	SaveUser(userID string, deviceID string, name string, desc string, age int, gender gender.Gender, hobbyTags []string, token string) (User, error)
}
