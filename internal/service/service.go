package service

import (
	"privacy-check/internal/repository"
)

type Service interface{}

type service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) Service {
	return &service{repo: repo}
}

func (s *service) Create() error {
	return nil
}
