package controllers

import (
	"github/murilorscampos/desafio-fabricio/database"
	"github/murilorscampos/desafio-fabricio/models"
	"github/murilorscampos/desafio-fabricio/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ConsultaContatosTelefone realiza a consulta dos contatos a partir de um número de teleefone informado
func ConsultaContatosTelefone(c *gin.Context) {

	telefone := models.Telefone{}

	telefoneID := c.Params.ByName("telefone")
	telefoneConvertido, _ := strconv.Atoi(telefoneID)

	if result := database.DB.Where(&models.Telefone{NumTelefone: telefoneConvertido}).Find(&telefone); result.Error != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"data": result.Error,
		})

		return

	} else if result.RowsAffected == 0 {

		c.JSON(http.StatusNotFound, gin.H{
			"data": "Contato não encontrado, telefone: " + telefoneID,
		})

		return

	}

	contato, err := BuscaClienteID(telefone.ContatoID)

	if err != nil {

		c.JSON(http.StatusNotFound, gin.H{
			"data": err.Error,
		})

		return
	}

	contatosComRecomendacao, err := utils.RealizaRecomendacao(contato)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"data": err.Error,
		})

	}

	c.JSON(http.StatusOK, contatosComRecomendacao)
}
