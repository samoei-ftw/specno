package models
type UpsertUser struct {
	UserID *uint `json:"user_id" validate:"required"` // * -> optional for insert
	Email        *string `json:"email"`
	Password *string `json:"password"`
}