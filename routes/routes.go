package routes

import (
	"github/murilorscampos/desafio-fabricio/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRoutes() {

	r := gin.Default()

	//contatos
	r.POST("/api/v1/contatos", controllers.InsereContato)
	r.PUT("/api/v1/contatos/id/:id", controllers.AlteraContato)
	r.DELETE("/api/v1/contatos/id/:id", controllers.ApagaContato)
	r.GET("/api/v1/contatos/", controllers.ListaContatos)
	r.GET("/api/v1/contatos/nome/:nome", controllers.ConsultaContatosNome)
	r.GET("/api/v1/contatos/email/:email", controllers.ConsultaContatosEmail)
	r.GET("/api/v1/contatos/cidade/:cidade", controllers.ConsultaContatosCidade)
	r.GET("/api/v1/contatos/uf/:uf", controllers.ConsultaContatosUF)
	r.GET("/api/v1/contatos/telefone/:telefone", controllers.ConsultaContatosTelefone)

	r.Run()

}
