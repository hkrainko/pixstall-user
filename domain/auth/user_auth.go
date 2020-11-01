package auth

type UserAuth interface {
	CreateNewToken(userID string, deviceID string) (string, error)
	Validate(tokenStr string) (*AppClaims, error)
}

type AppClaims struct {
	UserID    string
	DeviceID  string
	ExpiresAt int64
	IssuedAt  int64
}
