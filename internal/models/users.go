package models

// UserResponse represents the user model for API responses (without password)
type UserResponse struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}
