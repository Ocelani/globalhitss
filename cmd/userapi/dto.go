package main

import "globalhitss/pkg/user"

// User DTO.
type User struct {
	ID         uint   `gorm:"primaryKey"           json:"id,omitempty"`
	Nome       string `json:"nome,omitempty"`
	Sobrenome  string `json:"sobrenome,omitempty"`
	Contato    string `json:"contato,omitempty"`
	Endereco   string `json:"endereço,omitempty"`
	Nascimento string `json:"nascimento,omitempty"`
	CPF        string `json:"cpf,omitempty"`
}

// Response DTO.
type Response struct {
	User  User   `json:"user,omitempty"`
	Error string `json:"error,omitempty"`
}

// toUserEntity converts a User DTO to a user entity.
func toUserEntity(dto User) (entity *user.User) {
	return &user.User{
		ID:         user.ID(dto.ID),
		Nome:       dto.Nome,
		Sobrenome:  dto.Sobrenome,
		Contato:    dto.Contato,
		Endereco:   dto.Endereco,
		Nascimento: dto.Nascimento,
		CPF:        dto.CPF,
	}
}

// toUserDTO converts a user entity to a User DTO.
func toUserDTO(entity *user.User) (dto User) {
	return User{
		ID:         uint(entity.ID),
		Nome:       entity.Nome,
		Sobrenome:  entity.Sobrenome,
		Contato:    entity.Contato,
		Endereco:   entity.Endereco,
		Nascimento: entity.Nascimento,
		CPF:        entity.CPF,
	}
}

// toResponse converts a user entity to a Response DTO.
func toResponse(entity *user.User) (resp Response) {
	resp.User = toUserDTO(entity)
	return
}
