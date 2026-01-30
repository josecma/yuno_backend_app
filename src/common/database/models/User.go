package models

type User struct {
	ID          string
	FirstName   *string
	LastName    *string
	Username    string
	Password    string
	Email       string
	PhoneNumber string
}
