package repositories

import (
	"github.com/cyneptic/letsgo-authentication/internal/core/entities"
	"gorm.io/gorm"
)

// check user already register to prevent to register twice
func (p *Postgres) IsUserAlreadyRegisters(user entities.User) (int64, error) {
	res := p.db.Where("email = ?", user.Email).First(&user)
	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			return 0, nil // User not found, no error
		}
		return 0, res.Error // Other error occurred
	}
	return res.RowsAffected , nil
}

// add user to database ( registers user)
func (p *Postgres) AddUser(user entities.User) error {
	result := p.db.Create(user)
	return result.Error
}

func (p *Postgres) LoginHandler(email string) (*entities.User, error) {

	var fundedUser entities.User
	if err := p.db.Where("email = ? ", email).First(&fundedUser).Error; err != nil {
		return nil, err
	}
	return &fundedUser, nil
}