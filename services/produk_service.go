package services

import (
	"errors"
	"evermos/models"
	"evermos/repositories"
)

type ProdukInput struct {
	NamaProduk      string `json:"nama_produk"`
	Slug            string `json:"slug"`
	HargaReseller   int    `json:"harga_reseller"`
	HargaKonsumen   int    `json:"harga_konsumen"`
	Stok            int    `json:"stok"`
	Deskripsi       string `json:"deskripsi"`
	TokoID          uint   `json:"toko_id"`
	CategoryID      uint   `json:"category_id"`
}

type ProdukService interface {
	Create(input ProdukInput) (models.Produk, error)
	GetAll() ([]models.Produk, error)
	GetByID(id string) (models.Produk, error)
	Update(id string, input ProdukInput) (models.Produk, error)
	Delete(id string) error
}

type produkService struct {
	repo repositories.ProdukRepository
}

func NewProdukService(repo repositories.ProdukRepository) ProdukService {
	return &produkService{repo}
}

func (s *produkService) Create(input ProdukInput) (models.Produk, error) {
	produk := models.Produk{
		NamaProduk:    input.NamaProduk,
		Slug:          input.Slug,
		HargaReseller: input.HargaReseller,
		HargaKonsumen: input.HargaKonsumen,
		Stok:          input.Stok,
		Deskripsi:     input.Deskripsi,
		TokoID:        input.TokoID,
		CategoryID:    input.CategoryID,
	}
	return s.repo.Create(produk)
}

func (s *produkService) GetAll() ([]models.Produk, error) {
	return s.repo.FindAll()
}

func (s *produkService) GetByID(id string) (models.Produk, error) {
	return s.repo.FindByID(id)
}

func (s *produkService) Update(id string, input ProdukInput) (models.Produk, error) {
	exist, err := s.repo.FindByID(id)
	if err != nil {
		return exist, errors.New("produk not found")
	}
	exist.NamaProduk = input.NamaProduk
	exist.Slug = input.Slug
	exist.HargaReseller = input.HargaReseller
	exist.HargaKonsumen = input.HargaKonsumen
	exist.Stok = input.Stok
	exist.Deskripsi = input.Deskripsi
	exist.TokoID = input.TokoID
	exist.CategoryID = input.CategoryID

	return s.repo.Update(exist)
}

func (s *produkService) Delete(id string) error {
	_, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("produk not found")
	}
	return s.repo.Delete(id)
}
