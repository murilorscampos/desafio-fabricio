package utils

import (
	"encoding/json"
	"github/murilorscampos/desafio-fabricio/models"
	"io/ioutil"
	"log"
	"net/http"
	"unicode"
)

func IsMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}

// ConsultaClima realiza a consulta de clima da cidade utilizando a API da HGBrasil
func ConsultaClima(cidade, uf string) ([]byte, error) {

	const (
		chaveAPI  = "782b3097"
		camposAPI = "only_results,temp,condition_slug"
	)

	response, err := http.Get("https://api.hgbrasil.com/weather?fields=" + camposAPI + "&key=" + chaveAPI + "&city_name=" + cidade + "," + uf)

	if err != nil {

		log.Println("Erro ao consultar clima...", err.Error())

		return nil, err

	}

	dadosClima, err := ioutil.ReadAll(response.Body)

	if err != nil {

		log.Println("Erro ao converter retorno de API...", err.Error())

		return nil, err

	}

	return dadosClima, nil

}

// RealizaRecomendacao realiza a recomendacao de atividades dependendo do clima
func RealizaRecomendacao(contatos []models.Contato) ([]models.Contato, error) {

	for i := 0; i < len(contatos); i++ {

		var clima models.Clima

		dadosClima, err := ConsultaClima(contatos[i].Cidade, contatos[i].UF)

		if err != nil {

			return nil, err

		}

		json.Unmarshal(dadosClima, &clima)

		log.Println("Cidade:", contatos[i].Cidade, "UF:", contatos[i].UF, "Temperatura:", clima.Temperatura, "Condição:", clima.Condicao)

		recomendacao := geraRecomendacao(clima.Temperatura, clima.Condicao)

		contatos[i].Recomendacao = recomendacao

	}

	return contatos, nil
}

// geraRecomendacao verifica qual a recomendacao vai ser gerada dependendo clima
func geraRecomendacao(temperatura int, condicao string) string {

	var recomendacao string

	switch {
	case temperatura <= 18:
		recomendacao = "Gostaria de uma chocolate quente?"
	case temperatura >= 30:
		if condicao == "clear_day" || condicao == "clear_night" {
			recomendacao = "Vamos à praia?"
		} else if condicao == "rain" || condicao == "cloudly_day" || condicao == "cloudly_night" {
			recomendacao = "Vamos tomar uma sorvete?"
		}
	case temperatura > 18 && temperatura < 30:
		if condicao == "clear_day" || condicao == "clear_night" {
			recomendacao = "Vamos ir ao parque?"
		} else if condicao == "rain" || condicao == "cloudly_day" || condicao == "cloudly_night" {
			recomendacao = "Vamos ver um filme?"
		}
	default:
		recomendacao = "Sem recomendações por hoje..."
	}

	return recomendacao

}
