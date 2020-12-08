package domain

import "errors"

type Gender string
const (
	Male   Gender = "M"
	Female Gender = "F"
)

func New(s string) (Gender, error) {
	switch s {
	case string(Male):
		return Male, nil
	case string(Female):
		return Female, nil
	default:
		return "", errors.New("error")
	}
}