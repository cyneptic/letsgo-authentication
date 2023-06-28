package service

import (
	"errors"

	repositories "github.com/cyneptic/letsgo-authentication/infrastructure/repository"
	"github.com/cyneptic/letsgo-authentication/internal/core/entities"
	"github.com/cyneptic/letsgo-authentication/internal/core/ports"
)

type AuthenticationService struct {
	db ports.UserRepositoryContracts
}

func NewAuthenticationService() *AuthenticationService {
	db := repositories.NewPostgres()
	return &AuthenticationService{
		db: db,
	}
}
func (u *AuthenticationService) IsUserAlreadyRegisters(newUser entities.User) (bool , error) {
	res , err := u.db.IsUserAlreadyRegisters(newUser)

	if err != nil { 
		return false , err
	}

	if res > 0 {
		return true , nil
	}
	return false , nil
}
func (u *AuthenticationService) AddUser(newUser entities.User) error {

	isUserAlreadyExist , err := u.IsUserAlreadyRegisters(newUser)

	if err != nil { 
		return err
	}

	if isUserAlreadyExist == true {
		err := errors.New("User already registered")
		return err
	}

	err = u.db.AddUser(newUser)
	return err
}

func (u *AuthenticationService) LoginHandler(user entities.User) (string, error) {
	email := user.Email
	password := user.Password

	foundedUser, err := u.db.LoginHandler(email)

	if err != nil {
		return "", err
	}
	if foundedUser.Password != password {
		err := errors.New("email or password mismatch")
		return "", err
	}
	// token := GenerateToken(foundedUser.ID)
	
	// err = u.redis.AddToken(token)

	if err != nil {
		return "" , err
	}

	return "", nil
}