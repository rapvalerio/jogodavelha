package main

import "fmt"

func main() {
	var input string
	fmt.Println("_|_|_")
	fmt.Println("_|_|_")
	fmt.Println(" | | ")

	// fmt.Print("Digite algo: ")
	fmt.Scan(&input) // Lê a entrada do usuário e armazena em 'input'
	fmt.Println("Você digitou:", input)

}
