package service

import (
	"crypto/sha1"
	"fmt"
	"github.com/ariocp/go-app/internal/models"
	"github.com/ariocp/go-app/internal/repository"
	"github.com/dgrijalva/jwt-go"
	"log"
	"math/rand"
	"net/smtp"
	"os"
	"time"
)

const (
	tokenTTL = 12 * time.Hour
	codeTTL  = 15 * time.Minute
)

var (
	salt       = os.Getenv("SALT")
	signingKey = os.Getenv("SIGNING_KEY")
	smtpHost   = "smtp.rambler.ru"
	smtpPort   = "587"
	smtpUser   = "ariocp@rambler.ru"
	smtpPass   = "ImIs7QsZDI"
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
	user.ConfirmationCode = generateConfirmationCode()
	user.ConfirmationExpiry = time.Now().Add(codeTTL)
	user.IsConfirmed = false // пользователь помечается как неподтвержденный
	err := sendConfirmationEmail(user.Email, user.ConfirmationCode)
	if err != nil {
		return 0, err
	}
	return s.repo.CreateUser(user)
}

func (s *AuthService) ConfirmUser(username, code string) error {
	user, err := s.repo.GetUserByUsername(username)
	if err != nil {
		return err
	}

	if user.ConfirmationCode != code || time.Now().After(user.ConfirmationExpiry) {
		return fmt.Errorf("invalid or expired confirmation code")
	}

	return s.repo.ConfirmUser(username, code)
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

func generateConfirmationCode() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%04d", rand.Intn(10000))
}

func sendConfirmationEmail(to, code string) error {
	host := smtpHost
	port := smtpPort
	from := smtpUser
	pass := smtpPass

	auth := smtp.PlainAuth("", from, pass, host)
	msg := fmt.Sprintf("Subject: Confirmation Code\n\nYour confirmation code is: %s", code)

	err := smtp.SendMail(host+":"+port, auth, from, []string{to}, []byte(msg))
	if err != nil {
		log.Printf("smtp error: %s", err)
		return err
	}
	log.Print("Email sent successfully")
	return nil
}
