package model

type UserState string

const (
	UserStatePending    UserState = "P"
	UserStateActive     UserState = "A"
	UserStateTerminated UserState = "T"
)
