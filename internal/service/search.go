package service

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"privacy-check/internal/models"
)

func (s *service) SearchLeakDataById(userId int) (*models.LeakData, error) {
	leakData, err := s.repo.SearchUserLeakData(userId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			user, err := s.GetUserById(userId)
			if err != nil {
				return nil, err
			}

			dummyResponse, err := s.searchOnDummyJson(user.Email)
			if err != nil {
				if err.Error() == "not found" {
					return s.insertAndReturnLeakData(userId, models.DataStatusNotFound, nil)
				}
				return nil, fmt.Errorf("failed to search on DummyJson: %w", err)
			}

			return s.insertAndReturnLeakData(userId, models.DataStatusFound, dummyResponse)
		}
		return nil, err
	}

	return leakData, nil
}

func (s *service) insertAndReturnLeakData(userId int, status models.DataStatus, data []map[string]interface{}) (*models.LeakData, error) {
	newLeakData := &models.LeakData{
		UserID: userId,
		Status: status,
		Data:   data,
	}

	insertedId, err := s.repo.InsertUserLeakData(newLeakData)
	if err != nil {
		return nil, fmt.Errorf("failed to insert leak data into the database: %w", err)
	}

	newLeakData.ID = insertedId
	return newLeakData, nil
}

func (s *service) searchOnDummyJson(email string) ([]map[string]interface{}, error) {
	url := fmt.Sprintf("https://dummyjson.com/users/search?q=%s", email)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var dummyResponse models.DummyJsonResponse
	if err := json.Unmarshal(body, &dummyResponse); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	if dummyResponse.Total == 0 {
		return nil, errors.New("not found")
	}

	return dummyResponse.Users, nil
}
