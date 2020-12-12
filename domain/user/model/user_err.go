package model

import "fmt"

type UserError int

func (e UserError) Error() string {
	return fmt.Sprintf("%v", e)
}

const (
	UserErrorNotFound              UserError = 10
	UserErrorTerminated            UserError = 11
	UserErrorDuplicateUser         UserError = 12
	UserErrorAuthIDAlreadyRegister UserError = 13
	UserErrorUnknown               UserError = 99
)
