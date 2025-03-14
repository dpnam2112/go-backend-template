package dto

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

// CreateUserRequest represents the request payload for creating a user
type CreateUserRequest struct {
	Name string `json:"name" validate:"required"` // Name is required
}

// UserResponse represents the response format
type UserResponse struct {
	ID   uuid.UUID `json:"id"`   // UUID is required (automatically generated)
	Name string    `json:"name"` // Name is required
}

// Validate function for CreateUserRequest
func (c *CreateUserRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}
