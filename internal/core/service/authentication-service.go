package service

import (
	"errors"
	"log"

	repositories "github.com/cyneptic/letsgo-authentication/infrastructure/repository"
	"github.com/cyneptic/letsgo-authentication/internal/core/entities"
	"github.com/cyneptic/letsgo-authentication/internal/core/ports"
	
	"github.com/google/uuid"
)

type AuthenticationService struct {
	db    ports.UserRepositoryContracts
	redis ports.InMemoryRespositoryContracts
}

func NewAuthenticationService() *AuthenticationService {
	db := repositories.NewPostgres()
	redis := repositories.RedisInit()

	return &AuthenticationService{
		db:    db,
		redis: redis,
	}
}
func (u *AuthenticationService) IsUserAlreadyRegisters(newUser entities.User) (bool, error) {
	res, err := u.db.IsUserAlreadyRegisters(newUser)

	if err != nil {
		return true, err
	}

	if res > 0 {
		return true, nil
	}
	return false, nil
}
func (u *AuthenticationService) AddUser(newUser entities.User) error {

	isUserAlreadyExist, err := u.IsUserAlreadyRegisters(newUser)

	if err != nil {
		
		return err
	}

	if isUserAlreadyExist == true {
		err := errors.New("user already registered")
		return err
	}

	newUser.Password , _  = HashPassword(newUser.Password)

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

	decodedFoundedPassword := CheckPasswordHash(password , foundedUser.Password)

	if decodedFoundedPassword == false {
		err := errors.New("email or password mismatch")
		return "", err
	}
	token := GenerateToken(foundedUser.DBModel.ID)

	err = u.redis.AddToken(token)

	if err != nil {
		return "", err
	}

	return token, nil
}
func (u *AuthenticationService) AddToken(token string) {
	err := u.redis.AddToken(token)
	if err != nil {
		log.Fatal(err)
	}
}
func (u *AuthenticationService) Logout(token string) error {
	err := u.redis.RevokeToken(token)

	return err
}
func (u *AuthenticationService) TokenReceiver(token string) (string, error) {
	val, err := u.redis.TokenReceiver(token)
	return val, err
}
func (u *AuthenticationService) IsAdminAccount(id uuid.UUID) (bool, error) {
	isAdmin, err := u.db.IsAdminAccount(id)
	if err != nil {
		return false, err
	}
	return isAdmin, nil
}
func (u *AuthenticationService) Verify(number string, id uuid.UUID) (bool, error) {
	isVerified, err := u.db.Verify(number, id)
	if err != nil {
		return false, err
	}
	return isVerified, nil
}
