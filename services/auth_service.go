package services

import (
    "errors"
    "os"
    "time"
    "github.com/golang-jwt/jwt/v4"
    "golang.org/x/crypto/bcrypt"
    "gorm.io/gorm"
    "evermos/configs"
    "evermos/models"
)

type AuthService struct {
    DB *gorm.DB
}

func NewAuthService() *AuthService {
    return &AuthService{DB: configs.DB}
}

func (s *AuthService) Register(u *models.User) (*models.User, error) {
    hashed, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
    u.Password = string(hashed)
    if err := s.DB.Create(u).Error; err != nil {
        return nil, err
    }
    toko := models.Toko{
        UserID: u.ID,
        NamaToko: u.Nama + " Store",
    }
    if err := s.DB.Create(&toko).Error; err != nil {
        return nil, err
    }
    u.Toko = toko
    return u, nil
}

func (s *AuthService) Login(emailOrPhone, password string) (string, *models.User, error) {
    var user models.User
    if err := s.DB.Where("email = ? OR no_telp = ?", emailOrPhone, emailOrPhone).First(&user).Error; err != nil {
        return "", nil, errors.New("user not found")
    }
    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
        return "", nil, errors.New("invalid credentials")
    }
    secret := os.Getenv("JWT_SECRET")
    claims := jwt.MapClaims{
        "user_id": user.ID,
        "exp":     time.Now().Add(time.Hour * 72).Unix(),
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    t, err := token.SignedString([]byte(secret))
    if err != nil {
        return "", nil, err
    }
    return t, &user, nil
}
