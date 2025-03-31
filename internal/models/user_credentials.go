// UserCredentials is used for the login request body
package models

type UserCredentials struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
