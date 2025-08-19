package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	// nomePizzaria := "Pizzaria Go, "
	// precoPizza := 40.5
	// instagram := "@pizzariaGo"
	// telefone := 11904493
	// // := faz com que o código defina sozinho o tipo da variavel
	// fmt.Println(nomePizzaria, precoPizza, instagram, telefone)

	fmt.Println("Boas vindas ao jogo do número secreto")
	numeroMaximo := 1000
	rand.Seed(time.Now().UnixNano())
	numeroSecreto := rand.Intn(numeroMaximo) + 1
	fmt.Printf("Número secreto (para debug): %d\n", numeroSecreto)
	var chute int
	tentativas := 1

	for {
		fmt.Printf("Escolha um número entre 1 e %d: ", numeroMaximo)
		var input string
		fmt.Scanln(&input)
		chute, _ = strconv.Atoi(input)

		if chute == numeroSecreto {
			break
		} else {
			if chute > numeroSecreto {
				fmt.Printf("O número secreto é menor que %d\n", chute)
			} else {
				fmt.Printf("O número secreto é maior que %d\n", chute)
			}
			tentativas++
		}
	}

	palavraTentativa := "tentativa"
	if tentativas > 1 {
		palavraTentativa = "tentativas"
	}
	fmt.Printf("Isso aí! Você descobriu o número secreto %d com %d %s.\n", numeroSecreto, tentativas, palavraTentativa)
}
