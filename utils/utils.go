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

func RealizaRecomendacao(cidade, uf string) string {

	var (
		clima        models.Clima
		recomendacao string
	)

	dadosClima, err := consultaClima(cidade, uf)

	if err != nil {
		return ""
	}

	json.Unmarshal(dadosClima, &clima)

	log.Println("Cidade:", cidade, "UF:", uf, "Temperatura:", clima.Temperatura, "Descrição:", clima.Descricao)

	switch {
	case clima.Temperatura <= 18:
		recomendacao = "Gostaria de uma chocolate quente?"
	case clima.Temperatura >= 30:
		if clima.Descricao == "Tempo ensolarado" {
			recomendacao = "Vamos à praia?"
		} else if clima.Descricao == "Chuva" || clima.Descricao == "Chuvisco" {
			recomendacao = "Vamos tomar uma sorvete?"
		}
	case clima.Temperatura > 18 && clima.Temperatura < 30:
		if clima.Descricao == "Tempo ensolarado" {
			recomendacao = "Vamos ir ao parque?"
		} else if clima.Descricao == "Chuva" || clima.Descricao == "Chuvisco" {
			recomendacao = "Vamos ver um filme?"
		}
	default:
		recomendacao = "Sem recomendações por hoje..."
	}

	return recomendacao

}

func consultaClima(cidade, uf string) ([]byte, error) {

	const (
		chaveAPI  = "782b3097"
		camposAPI = "only_results,temp,description"
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
