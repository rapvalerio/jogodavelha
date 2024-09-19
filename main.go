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

//	type board struct {
//		board []*cube
//	}

func main() {
	var input string
	var tabuleiro = []int{
		4, 0, 0,
		0, 0, 0,
		0, 0, 0,
	}

	//TODO fazer as setas se moverem sem apertar o enter
	for {
		clearScreen()
		showBoard(tabuleiro)

		fmt.Print("Navegue usando -> ou <- e aperte o 'a' para marcar: ")

		fmt.Scan(&input)

		if input == "a" {
			addSymbol(tabuleiro)
		}

		if input == "\x1b[C" {
			fmt.Print("apertou a seta pra direita ")
			resetTabuleiro(tabuleiro, 1)
			// setNextFocus(tabuleiro)
		}

		if input == "\x1b[D" {
			fmt.Print("apertou a seta pra esquerda ")
			resetTabuleiro(tabuleiro, -1)
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
	for i := 0; i < len(tabuleiro); i++ {
		if tabuleiro[i] == 4 {
			tabuleiro[i] = 1
			setNextFocus(tabuleiro)
			break
		}
	}
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
