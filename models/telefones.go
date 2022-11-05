package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Telefone struct {
	gorm.Model  `json:"-"`
	ID          int `json:"id" gorm:"primaryKey,autoIncrement"`
	NumTelefone int `json:"numtelefone" validate:"nonzero"`
	ContatoID   int `json:"contatoid"`
}

func ValidaDadosTelefone(Telefone *Telefone) error {

	if err := validator.Validate(Telefone); err != nil {
		return err
	}

	return nil
}
