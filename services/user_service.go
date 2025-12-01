package services

import (
	"errors"
	"evermos/models"
	"evermos/repositories"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
	"github.com/golang-jwt/jwt/v4"
)

type RegisterInput struct {
	Nama  string `json:"nama"`
	Email string `json:"email"`
	Pass  string `json:"password"`
}

type LoginInput struct {
	Email string `json:"email"`
	Pass  string `json:"password"`
}

type UserService interface {
	Register(input RegisterInput) (models.User, error)
	Login(input LoginInput) (string, error)
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo}
}

func (s *userService) Register(input RegisterInput) (models.User, error) {
	hash, _ := bcrypt.GenerateFromPassword([]byte(input.Pass), 14)

	user := models.User{
		Nama:     input.Nama,
		Email:    input.Email,
		Password: string(hash),
	}

	return s.repo.Create(user)
}

func (s *userService) Login(input LoginInput) (string, error) {
	user, err := s.repo.FindByEmail(input.Email)
	if err != nil {
		return "", errors.New("email not found")
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Pass)) != nil {
		return "", errors.New("invalid password")
	}

	claims := jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := os.Getenv("JWT_SECRET")
	return token.SignedString([]byte(secret))
}
