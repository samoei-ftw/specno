package models

type ProjectUpdateRequest struct {
    Name        string `json:"name"`
    Description string `json:"description"`
}