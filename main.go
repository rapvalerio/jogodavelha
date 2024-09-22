package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/eiannone/keyboard"
)

const (
	playerX = 1
	playerO = 2
	focus   = 4
)

var currentPlayer = playerX

var combinacoesVencedoras = [][]int{
	{0, 1, 2}, {3, 4, 5}, {6, 7, 8},
	{0, 3, 6}, {1, 4, 7}, {2, 5, 8},
	{0, 4, 8}, {2, 4, 6},
}

var tabuleiro = []int{
	focus, 0, 0,
	0, 0, 0,
	0, 0, 0,
}

func main() {
	err := keyboard.Open()
	if err != nil {
		panic(err)
	}
	defer keyboard.Close()

	for {
		clearScreen()
		showInstructions()
		showBoard(tabuleiro)

		_, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}

		action(key)

		if key == keyboard.KeyEsc {
			break
		}

		if endgame() {
			break
		}
	}
}

func action(key keyboard.Key) {
	switch key {
	case keyboard.KeyEnter:
		addSymbol(tabuleiro)
	case keyboard.KeyArrowRight:
		resetTabuleiro(tabuleiro, 1)
	case keyboard.KeyArrowLeft:
		resetTabuleiro(tabuleiro, -1)
	case keyboard.KeyArrowDown:
		resetTabuleiro(tabuleiro, 3)
	case keyboard.KeyArrowUp:
		resetTabuleiro(tabuleiro, -3)
	}
}

func endgame() bool {
	if isWin(tabuleiro) {
		showBoard(tabuleiro)
		if currentPlayer == playerO {
			fmt.Printf("Jogador 2 venceu! \n")
		} else {
			fmt.Printf("Jogador 1 venceu! \n")
		}
		return true
	}

	if isDraw(tabuleiro) {
		showBoard(tabuleiro)
		fmt.Println("O jogo empatou")
		return true
	}

	return false
}

func showInstructions() {
	fmt.Println("Use as setas -> ou <- para navegar")
	fmt.Println("Aperta 'Enter' para marcar")
	fmt.Println("Aperta 'ESC' para sair")
	fmt.Println("")
}

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func showBoard(board []int) {
	clearScreen()
	showInstructions()
	for i := 0; i < 3; i++ {
		fmt.Printf("  %s | %s | %s \n", convertIntToShapes(board[i*3]), convertIntToShapes(board[i*3+1]), convertIntToShapes(board[i*3+2]))
		if i < 2 {
			fmt.Println(" ---+---+---")
		}
	}
}

func convertIntToShapes(valor int) string {
	switch valor {
	case playerO:
		return "O"
	case playerX:
		return "X"
	case focus:
		return "*"
	default:
		return " "
	}
}

func addSymbol(tabuleiro []int) {
	cursor := cursorPos(tabuleiro)
	if tabuleiro[cursor] == focus {
		tabuleiro[cursor] = currentPlayer
		togglePlayer()
		setNextFocus(tabuleiro)
	}
}

func isWin(tabuleiro []int) bool {
	for _, combinacoes := range combinacoesVencedoras {
		if tabuleiro[combinacoes[0]] != 0 && tabuleiro[combinacoes[0]] == tabuleiro[combinacoes[1]] && tabuleiro[combinacoes[1]] == tabuleiro[combinacoes[2]] {
			return true
		}
	}

	return false
}

func setNextFocus(tabuleiro []int) {
	for i := 0; i < len(tabuleiro); i++ {
		if tabuleiro[i] == 0 {
			tabuleiro[i] = focus
			break
		}
	}
}

func isDraw(tabuleiro []int) bool {
	for i := 0; i < len(tabuleiro); i++ {
		if tabuleiro[i] == 0 || tabuleiro[i] == focus {
			return false
		}
	}

	return true
}

func resetTabuleiro(tabuleiro []int, direcao int) []int {
	cursorPos := cursorPos(tabuleiro)

	if cursorPos == -1 {
		return tabuleiro
	}

	for i := 1; i < len(tabuleiro); i++ {
		nextPos := (cursorPos + direcao*i + len(tabuleiro)) % len(tabuleiro)
		if tabuleiro[nextPos] == 0 {
			tabuleiro[cursorPos] = 0
			tabuleiro[nextPos] = focus
			break
		}
	}

	return tabuleiro
}

func cursorPos(tabuleiro []int) int {
	cursorPos := -1

	for index, value := range tabuleiro {
		if value == focus {
			cursorPos = index
			break
		}
	}

	return cursorPos
}

func togglePlayer() {
	if currentPlayer == playerX {
		currentPlayer = playerO
	} else {
		currentPlayer = playerX
	}
}
