package ports

import (
	"github.com/cyneptic/letsgo-authentication/internal/core/entities"
	"github.com/google/uuid"
)

// اینترفیس دیتابیس هستش

type UserRepositoryContracts interface {
	
	AddUser(user entities.User) error
	Login(email string) (*entities.User, error)
	IsAdminAccount(id uuid.UUID) (bool, error)
	Verify(number string , id uuid.UUID) (bool, error)
	IsSuperAdminAccount(id uuid.UUID) (bool, error)
}

type InMemoryRespositoryContracts interface {
	AddToken(token string) error
	RevokeToken(token string) error
	TokenReceiver(token string) (string, error)
}
