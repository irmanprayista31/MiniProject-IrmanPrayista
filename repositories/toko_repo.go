package repositories

import (
	"evermos/models"
	"gorm.io/gorm"
)

type TokoRepository interface {
	Create(toko models.Toko) (models.Toko, error)
	FindAll() ([]models.Toko, error)
}

type tokoRepository struct {
	db *gorm.DB
}

func NewTokoRepository(db *gorm.DB) TokoRepository {
	return &tokoRepository{db}
}

func (r *tokoRepository) Create(toko models.Toko) (models.Toko, error) {
	err := r.db.Create(&toko).Error
	return toko, err
}

func (r *tokoRepository) FindAll() ([]models.Toko, error) {
	var toko []models.Toko
	err := r.db.Find(&toko).Error
	return toko, err
}
