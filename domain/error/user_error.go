package error

import "strconv"

type UserError int

func (e UserError) Error() string {
	switch e {
	case TerminatedUserError:
		return "TerminatedUserError"
	default:
		return strconv.Itoa(int(e))
	}
}

const (
	TerminatedUserError UserError = 301
)
