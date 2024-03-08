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
	return &UserService{
		Repository: repository,
	}
}

// Create data on PostgreSQL database.
func (s *UserService) Create(ctx context.Context, data *user.User) error {
	return s.Repository.Create(ctx, data)
}

// ReadOne user registered on the database.
func (s *UserService) ReadOne(ctx context.Context, id user.ID) (*user.User, error) {
	return s.Repository.ReadOne(ctx, id)
}

// Update data on PostgreSQL database.
func (s *UserService) Update(ctx context.Context, id user.ID, data *user.User) error {
	data.ID = id
	return s.Repository.Update(ctx, data)
}

// Delete data on PostgreSQL database.
func (s *UserService) Delete(ctx context.Context, id user.ID) error {
	return s.Repository.Delete(ctx, id)
}
