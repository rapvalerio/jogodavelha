package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/eiannone/keyboard"
)

// type cube struct {
// 	shape string
// 	color string
// 	focus int
// }

//	type board struct {
//		board []*cube
//	}
var player = 0

var combinacoesVencedoras = [][]int{
	{0, 1, 2}, {3, 4, 5}, {6, 7, 8},
	{0, 3, 6}, {1, 4, 7}, {2, 5, 8},
	{0, 4, 8}, {2, 4, 6},
}

var tabuleiro = []int{
	4, 0, 0,
	0, 0, 0,
	0, 0, 0,
}

func main() {
	// var input string

	//TODO fazer as setas se moverem sem apertar o enter
	for {
		clearScreen()
		showBoard(tabuleiro)

		err := keyboard.Open()
		if err != nil {
			panic(err)
		}
		defer keyboard.Close()

		fmt.Println("Navegue usando -> ou <- e aperte o 'a' para marcar")
		// fmt.Scan(&input)
		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}

		fmt.Printf("Você pressionou: %q (key code: %d)\n", char, key)

		if key == keyboard.KeyEnter {
			addSymbol(tabuleiro)
		}

		if key == keyboard.KeyArrowRight {
			fmt.Print("apertou a seta pra direita ")
			resetTabuleiro(tabuleiro, 1)
		}

		if key == keyboard.KeyArrowLeft {
			fmt.Print("apertou a seta pra esquerda ")
			resetTabuleiro(tabuleiro, -1)
		}

		if isWin(tabuleiro) {
			clearScreen()
			showBoard(tabuleiro)
			Jogador := player % 2
			fmt.Printf("Jogador %d venceu! \n", Jogador)
			break
		}
	}
}

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func showBoard(board []int) {
	fmt.Printf("  %s | %s | %s \n", convertIntToShapes(board[0]), convertIntToShapes(board[1]), convertIntToShapes(board[2]))
	fmt.Println(" ---+---+---")
	fmt.Printf("  %s | %s | %s \n", convertIntToShapes(board[3]), convertIntToShapes(board[4]), convertIntToShapes(board[5]))
	fmt.Println(" ---+---+---")
	fmt.Printf("  %s | %s | %s \n", convertIntToShapes(board[6]), convertIntToShapes(board[7]), convertIntToShapes(board[8]))
}

func convertIntToShapes(valor int) string {
	if valor == 2 {
		return "O"
	}

	if valor == 1 {
		return "X"
	}

	if valor == 4 {
		return "*"
	}

	return " "
}

func addSymbol(tabuleiro []int) {
	player++
	for i := 0; i < len(tabuleiro); i++ {
		if tabuleiro[i] == 4 {
			if player%2 == 0 {
				tabuleiro[i] = 2
			} else {
				tabuleiro[i] = 1
			}
			setNextFocus(tabuleiro)
			break
		}
	}
}

func isWin(tabuleiro []int) bool {
	// pedaco := combinacoesVencedoras[0]
	valor := []int{}
	for i := range combinacoesVencedoras {
		for _, k := range combinacoesVencedoras[i] {
			valor = append(valor, tabuleiro[k])
		}
		if valor[0] == valor[1] && valor[1] == valor[2] && valor[0] != 0 {
			return true
		}
		valor = []int{}
	}

	return false
}

func setNextFocus(tabuleiro []int) {
	for i := 0; i < len(tabuleiro); i++ {
		if tabuleiro[i] == 0 {
			tabuleiro[i] = 4
			break
		}
	}
}

func resetTabuleiro(tabuleiro []int, direcao int) []int {
	cursorPos := -1

	for i, v := range tabuleiro {
		if v == 4 {
			cursorPos = i
			break
		}
	}

	if cursorPos == -1 {
		return tabuleiro
	}

	for i := 1; i < len(tabuleiro); i++ {
		nextPos := (cursorPos + direcao*i + len(tabuleiro)) % len(tabuleiro)
		if tabuleiro[nextPos] == 0 {
			tabuleiro[cursorPos] = 0
			tabuleiro[nextPos] = 4
			break
		}
	}

	return tabuleiro
}
