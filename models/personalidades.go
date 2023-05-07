package models

type Personalidade struct {
	Nome     string `json:"nome"`
	Historia string `json:"historia"`
	Id       int    `json:"id"`
}

var Personalidades []Personalidade
