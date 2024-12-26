package service

import "privacy-check/internal/models"

func (s *service) Create(user *models.RegisterDTO) (int, error) {
	return s.repo.Create(&models.User{
		Firstname: user.FirstName,
		Lastname:  user.LastName,
		Email:     user.Email,
		Password:  generatePasswordHash(user.Password),
	})
}

func (s *service) GetUserById(userId int) (*models.User, error) {
	return s.repo.GetUserById(userId)
}
