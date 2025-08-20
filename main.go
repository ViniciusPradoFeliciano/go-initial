package main

import (
	"pizzaria/models"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/pizzas", getPizzas)

	// NomePizzaria := "Pizzaria Go, "
	// PrecoPizza := 40.5
	// instagram := "@pizzariaGo"
	// telefone := 11904493
	// // := faz com que o código defina sozinho o tipo da variavel

	// fmt.Println(NomePizzaria, PrecoPizza, instagram, telefone)
	// fmt.Println(pizzas)
	router.Run()
}

func getPizzas(c *gin.Context) {
	var pizzas = []models.Pizza{
		{ID: 1, Nome: "toscana", Preco: 40.50},
		{ID: 2, Nome: "mussarela", Preco: 35.50},
		{ID: 3, Nome: "4 queijos", Preco: 38.50},
	}
	c.JSON(200, gin.H{
		"pizzas": pizzas,
	})
}
