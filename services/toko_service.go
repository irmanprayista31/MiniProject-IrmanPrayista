package services

import (
	"evermos/models"
	"evermos/repositories"
)

type TokoInput struct {
	NamaToko string `json:"nama_toko"`
	UrlFoto  string `json:"url_foto"`
	UserID   uint   `json:"user_id"`
}

type TokoService interface {
	Create(input TokoInput) (models.Toko, error)
	GetAll() ([]models.Toko, error)
}

type tokoService struct {
	repo repositories.TokoRepository
}

func NewTokoService(repo repositories.TokoRepository) TokoService {
	return &tokoService{repo}
}

func (s *tokoService) Create(input TokoInput) (models.Toko, error) {
	toko := models.Toko{
		NamaToko: input.NamaToko,
		UrlFoto:  input.UrlFoto,
		UserID:   input.UserID,
	}

	return s.repo.Create(toko)
}

func (s *tokoService) GetAll() ([]models.Toko, error) {
	return s.repo.FindAll()
}
