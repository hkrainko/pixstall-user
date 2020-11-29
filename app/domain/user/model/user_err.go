package model

import "fmt"

type UserError int

func (e UserError) Error() string {
	return fmt.Sprintf("%v", e)
}

const (
	UserErrorNotFound UserError = 0
	UserErrorUnknown UserError = 1
)

