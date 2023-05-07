package main

import (
	"api-go-rest/database"
	"api-go-rest/models"
	"api-go-rest/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{"Nome 1", "Historia 1", 1},
		{"Nome 2", "Historia 2", 2},
	}
	database.ConectaComBancoDeDados()
	routes.HandleRequest()
}
