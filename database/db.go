package database

import (
	"github/murilorscampos/desafio-fabricio/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaBD() {

	// stringDeConexao := "host=localhost user=postgres password=root dbname=desafio-fabricio port=5432 sslmode=disable"
	stringDeConexao := "host=localhost user=desafio-fabricio password=postgres123456789 dbname=desafio-fabricio port=5000 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringDeConexao))

	if err != nil {
		log.Panic("Erro ao conectar no banco de dados.")
	} else {
		log.Println("Banco de dados conectado...")
	}

	DB.AutoMigrate(
		&models.Contato{},
		&models.Telefone{},
	)

}
