package ports

import (
	"github.com/cyneptic/letsgo-authentication/internal/core/entities"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

type UserServiceContract interface {
	
	AddUser(newUser entities.User) error
	LoginService(user entities.User) (string, error)
	Logout(token string) error
	IsAdminAccount(id uuid.UUID) (bool, error)
	Verify(number string , id uuid.UUID) (bool, error)
}

type InMemoryServiceContracts interface {
	AddToken(token string)
	RevokeToken(token string) *redis.StatusCmd
	TokenReceiver() (string, error)
}
