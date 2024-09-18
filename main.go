package main

import (
	"fmt"
	"os"
	"os/exec"
)

// type cube struct {
// 	shape string
// 	color string
// 	focus int
// }

// type board struct {
// 	board []*cube
// }
// var tabuleiro = []int{
// 	0,0,0,
// 	0,0,0,
// 	0,0,0,
// }

func main() {
	var input string
	focus := 0
	// shape := "X"

	//TODO fazer as setas se moverem sem apertar o enter
	//TODO como fazer o focus aparecer nas ultimas linhas
	for {
		clearScreen()
		showBoard(focus)

		fmt.Print("Escolha uma posição (1-9) ou 'q' para sair: ")

		fmt.Scan(&input)

		if input == "\x1b[A" {
			fmt.Print("apertou a seta pra cima ")
			if focus > 5 {
				focus = focus - 5
			}
		}

		if input == "\x1b[B" {
			fmt.Print("apertou a seta pra baixo ")
			if focus < 10 {
				focus = focus + 5
			}
		}

		if input == "\x1b[C" {
			fmt.Print("apertou a seta pra direita ")
			if focus == 4 {
				focus = 4
			} else {
				focus = focus + 2
			}
		}

		if input == "\x1b[D" {
			fmt.Print("apertou a seta pra esquerda ")
			if focus == 0 {
				focus = 0
			} else {
				focus = focus - 2
			}
		}

		fmt.Println("Você escolheu a posição:", input)
	}

	// fmt.Println("Jogo encerrado.")

}

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func showBoard(focus int) {
	var Reset = "\033[0m"
	// var Red = "\033[31m"
	// var White = "\033[97m"
	// var Green = "\033[32m"

	for i := 0; i < 15; i++ {
		var fundo = "\033[0m"
		// color := White

		if i == focus {
			// color = Red
			fundo = "\033[41m"
		}

		if i == 0 || i == 2 || i == 4 || i == 5 || i == 7 || i == 9 {
			if i == 4 || i == 9 {
				fmt.Println(fundo + "_" + Reset)
			} else {
				fmt.Print(fundo + "_" + Reset)
			}
		}

		if i == 10 || i == 12 || i == 14 {
			if i == 14 {
				fmt.Println(fundo + " " + Reset)
			} else {
				fmt.Print(fundo + " " + Reset)
			}
		}

		if i == 1 || i == 3 || i == 6 || i == 8 || i == 11 || i == 13 {
			if i == 13 {
				fmt.Println("|")
			} else {
				fmt.Print("|")
			}
		}

	}
}
