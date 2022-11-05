package main

import (
	"github/murilorscampos/desafio-fabricio/database"
	"github/murilorscampos/desafio-fabricio/routes"
	"log"
)

func main() {

	log.Println("Conectando ao banco de dados...")
	database.ConectaBD()

	log.Println("Iniciando servidor...")
	routes.HandleRoutes()
}
