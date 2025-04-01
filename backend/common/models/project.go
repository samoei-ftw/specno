package models

type Project struct {
	ID          uint   `json:"id"`
	UserId      uint   `json:"user_id"`
	Name        string `json:"name"`
	Description string `json:"descriptions"`
}
