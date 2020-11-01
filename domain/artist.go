package domain

type Artist struct {
	User `json:"user" validate:"required"`
}

type ArtistUsecase interface {
	FindByID(userID string) (*User, error)
	SaveUser(user User) error
}

type ArtistRepository interface {
	FindByID(userID string) (*User, error)
	SaveUser(user User) error
}
