package controllers

import (
	"github/murilorscampos/desafio-fabricio/database"
	"github/murilorscampos/desafio-fabricio/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ConsultaContatosTelefone(c *gin.Context) {

	telefones := []models.Telefone{}

	telefone := c.Params.ByName("telefone")
	telefoneConvertido, _ := strconv.Atoi(telefone)

	if result := database.DB.Where(&models.Telefone{NumTelefone: telefoneConvertido}).Find(&telefones).Order("contatoid"); result.Error != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"data": result.Error,
		})

		return

	} else if result.RowsAffected == 0 {

		c.JSON(http.StatusNotFound, gin.H{
			"data": "Contato não encontrado, telefone: " + telefone,
		})

		return

	}

	c.JSON(http.StatusOK, telefones)

}
