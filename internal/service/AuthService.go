package service

import (
	"Rent-And-Buy-App/internal/entity"
	"Rent-And-Buy-App/internal/repository"
	"Rent-And-Buy-App/pkg/auth"
	"Rent-And-Buy-App/pkg/hash"
	"errors"
)

type AuthService struct {
	repo       repository.UserRepository
	jwtManager *auth.JWTManager
}

type authService interface {
	Register(user *entity.User) error
	Login(email string, password string) (string, error)
}

func NewAuthService(repo repository.UserRepository, jwtManager *auth.JWTManager) *AuthService {
	return &AuthService{repo: repo, jwtManager: jwtManager}
}

func (as *AuthService) Register(user *entity.User) error {
	if exists, _ := as.repo.GetByEmail(user.Email); exists == nil {
		return errors.New("User with this email already exists")
	}

	hashed, err := hash.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashed

	err = as.repo.Create(user)
	return err
}

func (as *AuthService) Login(email string, password string) (string, error) {
	user, err := as.repo.GetByEmail(email)
	if err != nil || user == nil {
		return "", errors.New("This user does not exist")
	}

	isCorrect := hash.VerifyPassword(user.Password, password)
	if !isCorrect {
		return "", errors.New("Invalid password")
	}

	token, err := as.jwtManager.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}
