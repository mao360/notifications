package service

import (
	"context"
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/mao360/notifications/models"
	"github.com/mao360/notifications/pkg/repo"
)

const signature = "gd0394t389dfnvdsjnakjf23"
const salt = "82ghvsv89jven"

//go:generate go run github.com/vektra/mockery/v2@v2.30.1 --name=ServiceI
type ServiceI interface {
	NewUser(ctx context.Context, user *models.User) error
	GenerateToken(ctx context.Context, username, password string) (string, error)
	ParseToken(ctx context.Context, accessToken string) (*models.User, error)
	GetUser(ctx context.Context, username, password string) (bool, error)

	Subscribe(ctx context.Context, follower, author string) error
	Unsubscribe(ctx context.Context, follower, author string) error
	GetNotification(ctx context.Context, follower string) ([]string, error)
}

type Service struct {
	repo repo.RepoI
}

func NewService(repo repo.RepoI) *Service {
	return &Service{repo: repo}
}

func (s *Service) NewUser(ctx context.Context, user *models.User) error {
	user.Password = generatePasswordHash(user.Password)
	if user.UserName == "" || user.DateOfBirth == "" || user.Password == "" {
		return errors.New("not enough data for create user")
	}
	err := s.repo.NewUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GenerateToken(ctx context.Context, username, password string) (string, error) {
	user, err := s.repo.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", nil
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": user,
	})
	return token.SignedString([]byte(signature))
}
func (s *Service) ParseToken(ctx context.Context, accessToken string) (*models.User, error) {
	token, err := jwt.ParseWithClaims(accessToken, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		if method := token.Method.Alg(); method != "HS256" {
			return nil, errors.New("invalid sign in method")
		}
		return []byte(signature), nil
	})
	if err != nil {
		return nil, err
	}
	user, ok := token.Claims.(jwt.MapClaims)["user"]
	if !ok {
		return nil, errors.New("empty claims")
	}
	return user.(*models.User), nil
}

func (s *Service) GetUser(ctx context.Context, username, password string) (bool, error) {
	return s.repo.GetUser(username, generatePasswordHash(password))
}

func (s *Service) Subscribe(ctx context.Context, follower, author string) error {
	return s.repo.Subscribe(follower, author)
}
func (s *Service) Unsubscribe(ctx context.Context, follower, author string) error {
	return s.repo.Unsubscribe(follower, author)
}
func (s *Service) GetNotification(ctx context.Context, follower string) ([]string, error) {
	return s.repo.GetNotification(follower)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
