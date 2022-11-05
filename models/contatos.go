package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Contato struct {
	gorm.Model `json:"-"`
	ID         int        `json:"id" gorm:"primaryKey,autoIncrement"`
	Nome       string     `json:"nome" validate:"nonzero"`
	Logradouro string     `json:"logradouro"`
	Cidade     string     `json:"cidade"`
	Bairro     string     `json:"bairro"`
	UF         string     `json:"uf"`
	Email      string     `json:"mail"`
	Telefones  []Telefone `json:"telefones" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func ValidaDadosContato(Contato *Contato) error {

	if err := validator.Validate(Contato); err != nil {
		return err
	}

	return nil
}
