package service

import (
	"privacy-check/configs/env"
	"privacy-check/internal/models"
	"privacy-check/internal/repository"
)

type Service interface {
	GenerateToken(email, password string) (string, error)
	ParseToken(accessToken string) (int, error)

	Create(user *models.RegisterDTO) (int, error)
	GetUserById(userId int) (*models.User, error)

	SearchLeakDataById(userId int) (*models.LeakData, error)
}

type service struct {
	repo   repository.Repository
	config *env.EnvProject
}

func NewService(repo repository.Repository, config *env.EnvProject) Service {
	return &service{repo: repo, config: config}
}
