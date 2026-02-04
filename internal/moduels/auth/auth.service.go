package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/priyanshu334/go-bend/internal/moduels/user"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(email, password string) (*user.User, error)
	Login(email, password string) (string, string, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) Register(email, password string) (*user.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return nil, err
	}
	u := &user.User{
		Email:    email,
		Password: string(hash),
	}
	if err := s.repo.Create(u); err != nil {
		return nil, err
	}
	return u, nil
}

func (s *service) Login(email, password string) (string, string, error) {
	u, err := s.repo.FindByEmail(email)
	if err != nil {
		return "", "", errors.New("invalid email")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return "", "", errors.New("invalid credintials")
	}

	accessToken, err := generateToken(u.ID, 15*time.Hour)
	if err != nil {
		return "", "", err
	}
	refrestToken, err := generateToken(u.ID, 7*24*time.Hour)
	if err != nil {
		return "", "", err
	}
	return accessToken, refrestToken, nil
}

func generateToken(userID uint, ttl time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(ttl).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte("super-secret")) // move to env later
}
