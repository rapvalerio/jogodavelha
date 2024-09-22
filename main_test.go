package main

import "testing"

func TestTogglePlayer(t *testing.T) {
	currentPlayer = playerX
	togglePlayer()
	if currentPlayer != playerO {
		t.Errorf("Esperava jogar O, mas ficou jogador %d", currentPlayer)
	}

	togglePlayer()
	if currentPlayer != playerX {
		t.Errorf("Esperava jogar X, mas ficou jogador %d", currentPlayer)
	}
}

func TestIsDraw(t *testing.T) {
	tabuleiro = []int{
		1, 2, 2,
		2, 1, 1,
		2, 1, 1,
	}

	if !isDraw(tabuleiro) {
		t.Errorf("Não Deveria ter sido empate")
	}

	tabuleiro = []int{
		1, 2, 0,
		2, 4, 1,
		2, 1, 1,
	}

	if isDraw(tabuleiro) {
		t.Errorf("Deveria ter sido empate")
	}
}
