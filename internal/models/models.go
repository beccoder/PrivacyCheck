package models

type User struct {
	Id        int     `json:"id" db:"id" binding:"required"`
	Firstname string  `json:"firstname" db:"firstname" binding:"required"`
	Lastname  *string `json:"lastname" db:"lastname"`
	Email     string  `json:"email" db:"email" binding:"required"`
	Password  string  `json:"-" db:"password_hash" binding:"required"`
}
type DataStatus string

const (
	DataStatusFound    DataStatus = "FOUND"
	DataStatusNotFound DataStatus = "NOT_FOUND"
)

type LeakData struct {
	ID     int                      `json:"id" db:"id" binding:"required"`
	UserID int                      `json:"user_id" db:"user_id" binding:"required"`
	Status DataStatus               `json:"status" db:"status" binding:"required"`
	Data   []map[string]interface{} `json:"data,omitempty" db:"data,omitempty"`
}

type DummyJsonResponse struct {
	Users []map[string]interface{} `json:"users"`
	Total int                      `json:"total"`
	Skip  int                      `json:"skip"`
	Limit int                      `json:"limit"`
}
