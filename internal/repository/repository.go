package repository

import (
	"github.com/jmoiron/sqlx"
	"privacy-check/internal/models"
)

type Repository interface {
	Create(user *models.User) (int, error)
	GetUserByEmail(email string) (*models.User, error)
	GetUserById(id int) (*models.User, error)

	InsertUserLeakData(leakData *models.LeakData) (int, error)
	SearchUserLeakData(userId int) (*models.LeakData, error)
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{db: db}
}
