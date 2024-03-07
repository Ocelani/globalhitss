package userapi

import (
	"context"
	"globalhitss/pkg/user"
)

type UserService struct {
	Repository user.Repository
}

// NewUserService instantiates a new UserService object.
func NewUserService(repository user.Repository) *UserService {
	return &UserService{}
}

// Create data on PostgreSQL database.
func (r *UserService) Create(ctx context.Context, data *user.User) error {
	return nil
}

// ReadOne user registered on the database.
func (r *UserService) ReadOne(data *user.User) error {
	return nil
}

// Update data on PostgreSQL database.
func (r *UserService) Update(ctx context.Context, where, data *user.User) error {
	return nil
}

// Delete data on PostgreSQL database.
func (r *UserService) Delete(ctx context.Context, data *user.User) error {
	return nil
}
