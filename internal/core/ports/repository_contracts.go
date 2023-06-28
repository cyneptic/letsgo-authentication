package ports

import (
	"github.com/cyneptic/letsgo-authentication/internal/core/entities"
	

)

// اینترفیس دیتابیس هستش

type UserRepositoryContracts interface {
	IsUserAlreadyRegisters(user entities.User) (int64 , error)
	AddUser(user entities.User) error
	LoginHandler(email string) (*entities.User, error)
}