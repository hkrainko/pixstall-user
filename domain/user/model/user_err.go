package model

import "strconv"

type UserError int

func (e UserError) Error() string {
	return strconv.Itoa(int(e))
}

const (
	UserErrorNotFound              UserError = 10
	UserErrorTerminated            UserError = 11
	UserErrorDuplicateUser         UserError = 12
	UserErrorAuthIDAlreadyRegister UserError = 13
	UserErrorUnknown               UserError = 99
)
