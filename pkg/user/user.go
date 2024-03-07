package user

import (
	"context"
)

type ID uint

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
	ReadOne(context.Context, *User) error
	Update(ctx context.Context, where, update *User) error
	Delete(context.Context, *User) error
}

// Service is the interface that wraps the methods of the User service.
type Service interface {
	Register(context.Context, *User) error
	FindOne(context.Context, *User) error
	Update(ctx context.Context, where, update *User) error
	Remove(context.Context, *User) error
}
