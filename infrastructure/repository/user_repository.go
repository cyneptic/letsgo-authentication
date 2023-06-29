package repositories

import (
	"context"
	"errors"
	"time"

	"github.com/cyneptic/letsgo-authentication/internal/core/entities"
	"github.com/google/uuid"
	
	"gorm.io/gorm"
)

// check user already register to prevent to register twice
func (p *Postgres) IsUserAlreadyRegisters(user entities.User) (int64, error) {
	var count int64
	res := p.db.Model(entities.User{}).Where("email = ?", user.Email).Count(&count)
	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			return 0, nil // User not found, no error
		}
		print("auth respository 16", res.Error)
		return 0, res.Error // Other error occurred
	}
	return count, nil
}

// add user to database ( registers user)
func (p *Postgres) AddUser(user entities.User) error {
	result := p.db.Create(&user)
	return result.Error
}

func (p *Postgres) LoginHandler(email string) (*entities.User, error) {

	var fundedUser entities.User
	if err := p.db.Where("email = ? ", email).First(&fundedUser).Error; err != nil {
		return nil, err
	}
	return &fundedUser, nil
}

// redis
func (r *RedisDB) AddToken(token string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := r.client.Set(ctx, token, true, 0).Err()
	if err != nil {
		panic(err)
	}
	return err

}
func (r *RedisDB) RevokeToken(token string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := r.client.Set(ctx, token, false, 0).Err()
	return err
}

func (r *RedisDB) TokenReceiver(token string) (string, error) {
	ctx := context.Background()
	val, err := r.client.Get(ctx, token).Result()

	return val, err
}

func (p *Postgres) IsAdminAccount(id uuid.UUID) (bool, error) {
	var user entities.User
	res := p.db.Where("id = ?", id).First(&user)
	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			err := errors.New("there is no account with this ID")
			return false, err
		}
		return false, res.Error
	}
	if user.Role == "admin" {
		return true, nil
	}
	return false, nil
}

func (p *Postgres) Verify(number string , id uuid.UUID) (bool, error) {
	var user entities.User
	res := p.db.Where("id = ?", id).First(&user)
	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			err := errors.New("there is no account with this ID")
			return false, err
		}
		return false, res.Error
	}
	if user.PhoneNumber != number { 
		return false, nil
	}
	return true, nil
}