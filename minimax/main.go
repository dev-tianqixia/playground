package main

import (
	"github.com/dev-tianqixia/playground/minimax/game"
	"math"
)

// simulates a tic-tac-toe game with minimax algorithm
func main() {
	game.NewGame(game.NewBoard(3), math.MaxInt32).Run()
}
