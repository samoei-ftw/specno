package models

type Project struct {
	ID          int    `json:"id"`
	UserID      uint    `json:"user_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
