// UserCredentials is used for the login request body
package models

type UserCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
