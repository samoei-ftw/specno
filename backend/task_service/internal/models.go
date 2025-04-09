package internal
var taskCreateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	UserID      uint    `json:"user_id"`
}