package models

type Clima struct {
	Temperatura int    `json:"temp"`
	Condicao    string `json:"condition_slug"`
}
