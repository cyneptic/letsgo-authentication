package ports

import "github.com/cyneptic/letsgo-authentication/internal/core/entities"

type UserServiceContract interface {
	IsUserAlreadyRegisters(newUser entities.User) (bool, error)
	AddUser(newUser entities.User) error
	LoginHandler(user entities.User) (string, error)
}