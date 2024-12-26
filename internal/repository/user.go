package repository

import (
	"encoding/json"
	"fmt"
	"privacy-check/internal/models"
)

func (r *repository) Create(user *models.User) (int, error) {
	var (
		id    int
		query = `
			INSERT INTO users 
    			(firstname, lastname, email, password_hash) 
			VALUES ($1, $2, $3, $4) 
			RETURNING id
`
	)

	row := r.db.QueryRow(query, user.Firstname, user.Lastname, user.Email, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repository) GetUserByEmail(email string) (*models.User, error) {
	var (
		user  models.User
		query = `
			SELECT id, firstname, lastname, email, password_hash 
			FROM users 
			WHERE email=$1
`
	)

	err := r.db.Get(&user, query, email)

	return &user, err
}

func (r *repository) GetUserById(id int) (*models.User, error) {
	var (
		user  models.User
		query = `
			SELECT id, firstname, lastname, email, password_hash 
			FROM users 
			WHERE id=$1
`
	)

	err := r.db.Get(&user, query, id)

	return &user, err
}

func (r *repository) InsertUserLeakData(leakData *models.LeakData) (int, error) {
	dataJSON, err := json.Marshal(leakData.Data)
	if err != nil {
		return 0, fmt.Errorf("failed to marshal data to JSON: %w", err)
	}

	var (
		id    int
		query = `
			INSERT INTO leak_data 
    			(user_id, status, data) 
			VALUES ($1, $2, $3) 
			RETURNING id
`
	)

	row := r.db.QueryRow(query, leakData.UserID, leakData.Status, dataJSON)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repository) SearchUserLeakData(userId int) (*models.LeakData, error) {
	var (
		leakData models.LeakData
		query    = `
			SELECT id, user_id, status, data 
			FROM leak_data 
			WHERE user_id=$1
`
	)

	var rawData []byte
	err := r.db.QueryRow(query, userId).Scan(&leakData.ID, &leakData.UserID, &leakData.Status, &rawData)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(rawData, &leakData.Data)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSONB data: %w", err)
	}

	return &leakData, nil
}
