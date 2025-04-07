package models

type Project struct {
	ID          uint   `json:"id"`
	UserID 		uint   `json:"user_id"`
	Name        string `json:"name"`
	Description string `json:"descriptions"`
}
