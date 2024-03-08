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
	Endereço   string
	Nascimento string
	CPF        string
}

// Repository is the interface that wraps the methods of the User repository.
type Repository interface {
	Create(context.Context, *User) error
	ReadOne(context.Context, ID) (*User, error)
	Update(context.Context, *User) error
	Delete(context.Context, ID) error
}

// Service is the interface that wraps the methods of the User service.
type Service interface {
	Register(context.Context, *User) error
	FindOne(context.Context, ID) (*User, error)
	Update(context.Context, ID, *User) error
	Remove(context.Context, ID) error
}
