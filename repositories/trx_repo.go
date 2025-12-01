package repositories

import (
	"evermos/models"
	"gorm.io/gorm"
)

type TrxRepository interface {
	Create(input models.Trx) (models.Trx, error)
}

type trxRepo struct {
	db *gorm.DB
}

func NewTrxRepo(db *gorm.DB) TrxRepository {
	return &trxRepo{db}
}

func (r *trxRepo) Create(input models.Trx) (models.Trx, error) {
	err := r.db.Create(&input).Error
	return input, err
}
