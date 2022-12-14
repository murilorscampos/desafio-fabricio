package controllers

import (
	"github/murilorscampos/desafio-fabricio/database"
	"github/murilorscampos/desafio-fabricio/models"
	"github/murilorscampos/desafio-fabricio/utils"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"gorm.io/gorm"
)

// InsereContato realiza a inclusao de um novo contato e seus telefones
func InsereContato(c *gin.Context) {

	contato := models.Contato{}

	if result := c.ShouldBindJSON(&contato); result != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"data": result.Error,
		})

		return

	}

	if result := models.ValidaDadosContato(&contato); result != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"data": result.Error,
		})

		return

	}

	if result := database.DB.Create(&contato); result.Error != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"data": result.Error,
		})

		return

	}

	c.JSON(http.StatusOK, gin.H{
		"data": "Contato inserido",
	})

}

// AlteraContato realiza alteracoes nas informacoes de contato e telefone
func AlteraContato(c *gin.Context) {

	contato := models.Contato{}

	id := c.Params.ByName("id")

	if result := database.DB.Model(&contato).Preload("Telefones").First(&contato, id).RowsAffected; result == 0 {

		c.JSON(http.StatusNotFound, gin.H{
			"data": "Contato não encontrado, ID: " + id,
		})

		return

	}

	if result := c.ShouldBindJSON(&contato); result != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"data": result.Error,
		})

		return

	}

	if result := models.ValidaDadosContato(&contato); result != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"data": result.Error,
		})

		return

	}

	if result := database.DB.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&contato); result.Error != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"data": result.Error,
		})

		return

	}

	c.JSON(http.StatusOK, gin.H{
		"data": "Contato atualizado, ID: " + id,
	})

}

// ApagaContato realiza a exclusao do contato e dos telefones associados
func ApagaContato(c *gin.Context) {

	id := c.Params.ByName("id")
	idConvertido, _ := strconv.Atoi(id)

	if result := database.DB.Select("Telefones").Delete(&models.Contato{ID: idConvertido}); result.Error != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"data": result.Error,
		})

		return

	} else if result.RowsAffected == 0 {

		c.JSON(http.StatusNotFound, gin.H{
			"data": "Contato não encontrado, ID: " + id,
		})

		return

	}

	c.JSON(http.StatusOK, gin.H{
		"data": "Contato excluído, ID: " + id,
	})

}

// ListaContatos exibe todos os contatos
func ListaContatos(c *gin.Context) {

	contatos := []models.Contato{}

	if result := database.DB.Preload("Telefones").Find(&contatos); result.Error != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"data": result.Error,
		})

		return

	}

	contatosComRecomendacao, err := utils.RealizaRecomendacao(contatos)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"data": err.Error,
		})

	}

	c.JSON(http.StatusOK, contatosComRecomendacao)

}

// ConsultaContatosNome realiza a consulta dos contatos com base em um nome informado
func ConsultaContatosNome(c *gin.Context) {

	contatos := []models.Contato{}

	nome := c.Params.ByName("nome")

	t := transform.Chain(norm.NFD, transform.RemoveFunc(utils.IsMn), norm.NFC)

	nomeSemAcento, _, _ := transform.String(t, nome)

	if result := database.DB.Preload("Telefones").Where("lower(unaccent(nome)) LIKE ?", strings.ToLower(nomeSemAcento)+"%").Find(&contatos).Order("nome"); result.Error != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"data": result.Error,
		})

		return

	} else if result.RowsAffected == 0 {

		c.JSON(http.StatusNotFound, gin.H{
			"data": "Contato não encontrado, nome: " + nome,
		})

		return

	}

	contatosComRecomendacao, err := utils.RealizaRecomendacao(contatos)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"data": err.Error,
		})

	}

	c.JSON(http.StatusOK, contatosComRecomendacao)

}

// ConsultaContatosEmail realiza a consulta dos contatos com base em um e-mail informado
func ConsultaContatosEmail(c *gin.Context) {

	contatos := []models.Contato{}

	email := c.Params.ByName("email")

	if result := database.DB.Preload("Telefones").Where("lower(email) LIKE ?", strings.ToLower(email)+"%").Find(&contatos).Order("nome"); result.Error != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"data": result.Error,
		})

		return

	} else if result.RowsAffected == 0 {

		c.JSON(http.StatusNotFound, gin.H{
			"data": "Contato não encontrado, e-mail: " + email,
		})

		return

	}

	contatosComRecomendacao, err := utils.RealizaRecomendacao(contatos)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"data": err.Error,
		})

	}

	c.JSON(http.StatusOK, contatosComRecomendacao)

}

// ConsultaContatosCidade realiza a consulta dos contatos com base em um cidade informada
func ConsultaContatosCidade(c *gin.Context) {

	contatos := []models.Contato{}

	cidade := c.Params.ByName("cidade")

	t := transform.Chain(norm.NFD, transform.RemoveFunc(utils.IsMn), norm.NFC)

	cidadeSemAcento, _, _ := transform.String(t, cidade)

	if result := database.DB.Preload("Telefones").Where("lower(unaccent(cidade)) LIKE ?", strings.ToLower(cidadeSemAcento)+"%").Find(&contatos).Order("nome"); result.Error != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"data": result.Error,
		})

		return

	} else if result.RowsAffected == 0 {

		c.JSON(http.StatusNotFound, gin.H{
			"data": "Contato não encontrado, cidade: " + cidade,
		})

		return

	}

	contatosComRecomendacao, err := utils.RealizaRecomendacao(contatos)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"data": err.Error,
		})

	}

	c.JSON(http.StatusOK, contatosComRecomendacao)

}

// ConsultaContatosUF realiza a consulta dos contatos com base em um UF informada
func ConsultaContatosUF(c *gin.Context) {

	contatos := []models.Contato{}

	uf := c.Params.ByName("uf")

	if result := database.DB.Preload("Telefones").Where("lower(uf) LIKE ?", strings.ToLower(uf)+"%").Find(&contatos).Order("nome"); result.Error != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"data": result.Error,
		})

		return

	} else if result.RowsAffected == 0 {

		c.JSON(http.StatusNotFound, gin.H{
			"data": "Contato não encontrado, UF: " + uf,
		})

		return

	}

	contatosComRecomendacao, err := utils.RealizaRecomendacao(contatos)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"data": err.Error,
		})

	}

	c.JSON(http.StatusOK, contatosComRecomendacao)

}

func BuscaClienteID(clienteID int) ([]models.Contato, error) {

	contatos := []models.Contato{}

	if result := database.DB.Preload("Telefones").Find(&contatos, clienteID); result.Error != nil {

		return contatos, result.Error

	}

	return contatos, nil

}
