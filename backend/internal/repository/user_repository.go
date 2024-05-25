package repository

import "main/pkg/types"

// UserRepository defines the methods a user repository must implement.
type UserRepository interface {
	GetAllUsers() ([]types.User, error)
	CreateUser(user *types.User) error
	// GetUserByID(userID int) (*types.User, error)
}

