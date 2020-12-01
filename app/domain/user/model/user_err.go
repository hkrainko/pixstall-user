package model

import "fmt"

type UserError int

func (e UserError) Error() string {
	return fmt.Sprintf("%v", e)
}

const (
	UserErrorNotFound UserError = 0
	UserErrorInactive UserError = 1
	UserErrorUnknown UserError = 99
)

