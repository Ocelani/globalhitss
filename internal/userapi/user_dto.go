package userapi

type User struct {
	ID         uint   `gorm:"primaryKey" json:"id,omitempty"`
	Nome       string `json:"nome,omitempty"`
	Sobrenome  string `json:"sobrenome,omitempty"`
	Contato    string `json:"contato,omitempty"`
	Endereço   string `json:"endereço,omitempty"`
	Nascimento string `json:"nascimento,omitempty"`
	CPF        string `json:"cpf,omitempty"`
}
