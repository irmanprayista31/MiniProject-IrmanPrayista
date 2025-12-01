package repositories

import (
	"evermos/models"
	"gorm.io/gorm"
)

type ProdukRepository interface {
	Create(input models.Produk) (models.Produk, error)
	FindAll() ([]models.Produk, error)
	FindByID(id string) (models.Produk, error)
	Update(input models.Produk) (models.Produk, error)
	Delete(id string) error
}

type produkRepo struct {
	db *gorm.DB
}

func NewProdukRepo(db *gorm.DB) ProdukRepository {
	return &produkRepo{db}
}

func (r *produkRepo) Create(input models.Produk) (models.Produk, error) {
	err := r.db.Create(&input).Error
	return input, err
}

func (r *produkRepo) FindAll() ([]models.Produk, error) {
	var data []models.Produk
	err := r.db.Find(&data).Error
	return data, err
}

func (r *produkRepo) FindByID(id string) (models.Produk, error) {
	var produk models.Produk
	err := r.db.First(&produk, "id = ?", id).Error
	return produk, err
}

func (r *produkRepo) Update(input models.Produk) (models.Produk, error) {
	err := r.db.Save(&input).Error
	return input, err
}

func (r *produkRepo) Delete(id string) error {
	return r.db.Delete(&models.Produk{}, "id = ?", id).Error
}
