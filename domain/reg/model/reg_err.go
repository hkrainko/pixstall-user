package model

type RegError int

func (e RegError) Error() string {
	switch e {
	case RegErrorDuplicateUser:
		return "RegErrorDuplicateUser"
	case RegErrorAuthIDAlreadyRegister:
		return "RegErrorAuthIDAlreadyRegister"
	default:
		return "RegErrorUnknown"
	}
}

const (
	RegErrorDuplicateUser         RegError = 10
	RegErrorAuthIDAlreadyRegister RegError = 11
	RegErrorUnknown               RegError = 99
)
