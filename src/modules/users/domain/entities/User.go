package entities

import (
	"errors"
	"strings"
)

type User struct {
	ID          string
	FirstName   string
	LastName    string
	Username    string
	Password    string
	Email       string
	PhoneNumber string
}

func (*User) NewUser(ID, firstName, lastName, username, email, password string) (*User, error) {

	user := &User{
		ID: ID,
	}

	if err := user.SetFirstName(firstName); err != nil {
		return nil, err
	}

	if err := user.SetLastName(lastName); err != nil {
		return nil, err
	}

	if err := user.SetUsername(username); err != nil {
		return nil, err
	}

	if err := user.SetEmail(email); err != nil {
		return nil, err
	}

	if err := user.SetPassword(password); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *User) SetID(id string) error {
	if id == "" {
		return errors.New("ID cannot be empty")
	}
	u.ID = id
	return nil
}

func (u *User) SetFirstName(firstName string) error {
	if firstName == "" {
		return errors.New("first name cannot be empty")
	}
	u.FirstName = strings.TrimSpace(firstName)
	return nil
}

func (u *User) SetLastName(lastName string) error {
	if lastName == "" {
		return errors.New("last name cannot be empty")
	}
	u.LastName = strings.TrimSpace(lastName)
	return nil
}

func (u *User) SetUsername(username string) error {
	if username == "" {
		return errors.New("username cannot be empty")
	}
	u.Username = strings.TrimSpace(username)
	return nil
}

func (u *User) SetPassword(password string) error {
	if password == "" {
		return errors.New("password cannot be empty")
	}
	u.Password = password
	return nil
}

func (u *User) SetEmail(email string) error {
	if email == "" {
		return errors.New("email cannot be empty")
	}
	u.Email = strings.TrimSpace(strings.ToLower(email))
	return nil
}

func (u *User) SetPhoneNumber(phoneNumber string) error {

	if phoneNumber == "" {
		u.PhoneNumber = ""
		return nil
	}
	u.PhoneNumber = strings.TrimSpace(phoneNumber)
	return nil
}
