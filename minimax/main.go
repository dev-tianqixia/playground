package main

import (
	"github.com/dev-tianqixia/playground/minimax/game"
)

// simulates a tic-tac-toe game with minimax algorithm
func main() {
	// game.NewGame(game.NewBoard(3), math.MaxInt32).Run()
	game.NewGame(game.NewBoard(4), 5).Run()
}
