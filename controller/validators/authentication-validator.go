package validators

import (
	"errors"

	"github.com/cyneptic/letsgo-authentication/internal/core/entities"
)

func LoginValidation(u entities.User) error {
	if u.Email == "" {
		err := errors.New("please enter a valid email")
		return err
	}
	if u.Password == "" {
		err := errors.New("please enter a valid password")
		return err
	}
	return nil
}

func LogoutValidation(h string) error {
	if h == "" {
		err := errors.New("please enter a valid token")
		return err
	}
	return nil
}

func RegisterValidation(u entities.User) error {
	if u.Name == "" {
		err := errors.New("please enter a valid name")
		return err
	}
	if u.DateOfBirth.IsZero() {
		err := errors.New("please enter a valid name")
		return err
	}
	if u.PhoneNumber == "" {
		err := errors.New("please enter a valid phone number")
		return err
	}
	if u.Email == "" {
		err := errors.New("please enter a valid email")
		return err
	}
	if u.Password == "" {
		err := errors.New("please enter a valid password")
		return err
	}
	return nil
}
func IsAdmin(id string) error {
	if id == "" {
		err := errors.New("please enter a valid id")
		return err
	}
	return nil
}
func VerifyValidation(number, id string) error {
	if number == "" || id == "" {
		err := errors.New("please enter a valid id and phone number")
		return err
	}
	return nil
}
