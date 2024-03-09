package user

import (
	"context"
)

// ID is the type of the user ID.
type ID uint

// User is the model of the user entity.
type User struct {
	ID         ID
	Nome       string
	Sobrenome  string
	Contato    string
	Endereco   string
	Nascimento string
	CPF        string
}

// Repository is the interface that wraps the methods of the User repository.
type Repository interface {
	Create(ctx context.Context, data *User) error
	ReadOne(ctx context.Context, id ID) (*User, error)
	Update(ctx context.Context, data *User) error
	Delete(ctx context.Context, id ID) error
}

// Service is the interface that wraps the methods of the User service.
type Service interface {
	Register(ctx context.Context, data *User) error
	FindOne(ctx context.Context, id ID) (*User, error)
	Update(ctx context.Context, id ID, data *User) error
	Remove(ctx context.Context, id ID) error
}
