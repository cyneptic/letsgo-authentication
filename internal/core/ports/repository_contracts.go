package ports

import (
	"github.com/cyneptic/letsgo-authentication/internal/core/entities"
	"github.com/google/uuid"
)

// اینترفیس دیتابیس هستش

type UserRepositoryContracts interface {
	IsUserAlreadyRegisters(user entities.User) (int64 , error)
	AddUser(user entities.User) error
	LoginHandler(email string) (*entities.User, error)
	IsAdminAccount(id uuid.UUID) (bool, error)
	Verify(number string , id uuid.UUID) (bool, error)
}

type InMemoryRespositoryContracts interface {
	AddToken(token string) error
	RevokeToken(token string) error
	TokenReceiver(token string) (string, error)
}
