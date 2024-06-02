package service

import (
	"crypto/sha1"
	"fmt"
	"github.com/ariocp/go-app/internal/models"
	"github.com/ariocp/go-app/internal/repository"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

const tokenTTL = 12 * time.Hour

var (
	salt       = os.Getenv("SALT")
	signingKey = os.Getenv("SIGNING_KEY")
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int64 `json:"user_id"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user models.User) (int64, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})
	return token.SignedString([]byte(signingKey))
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
