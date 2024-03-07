package userapi

import (
	"context"
	"globalhitss/pkg/infra/database"
	"globalhitss/pkg/user"
)

// UserRepository type uses a Postgres database to sync stored data.
// Implements: user.Repository.
type UserRepository struct{ DB *database.Postgres }

// NewUserRepository instantiates a new UserRepository object.
func NewUserRepository(db *database.Postgres) *UserRepository {
	u := &User{}
	if err := db.Model(u).AutoMigrate(u); err != nil {
		panic(err)
	}
	return &UserRepository{db}
}

// Create data on PostgreSQL database.
func (r *UserRepository) Create(ctx context.Context, data *user.User) error {
	return r.DB.Create(data).Error
}

// ReadOne user registered on the database.
func (r *UserRepository) ReadOne(ctx context.Context, data *user.User) error {
	return r.DB.First(data).Error
}

// Update data on PostgreSQL database.
func (r *UserRepository) Update(ctx context.Context, where, data *user.User) error {
	return r.DB.Model(where).Updates(data).Error
}

// Delete data on PostgreSQL database.
func (r *UserRepository) Delete(ctx context.Context, data *user.User) error {
	return r.DB.Delete(data).Error
}