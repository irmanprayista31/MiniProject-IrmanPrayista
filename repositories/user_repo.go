package repositories

import (
	"evermos/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(input models.User) (models.User, error)
	FindByEmail(email string) (models.User, error)
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepository {
	return &userRepo{db}
}

func (r *userRepo) Create(input models.User) (models.User, error) {
	err := r.db.Create(&input).Error
	return input, err
}

func (r *userRepo) FindByEmail(email string) (models.User, error) {
	var user models.User
	err := r.db.First(&user, "email = ?", email).Error
	return user, err
}
