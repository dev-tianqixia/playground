package game

import (
	"fmt"
	"math"
)

const threeInARow = 3

type Game struct {
	board           *Board
	self            player
	opponent        player
	winner          player
	round           int
	simulationDepth int // unused
}

func NewGame(board *Board, simulationDepth int) *Game {
	return &Game{
		board:           board,
		simulationDepth: simulationDepth,
	}
}

func (g *Game) Run() {
	currentPlayer := O
	for {
		winner, terminated := g.board.IsTerminated(threeInARow)
		if terminated {
			fmt.Println("game over, winner is: ", winner)
			return
		}

		g.round += 1
		g.self, g.opponent = currentPlayer, turn(currentPlayer)
		fmt.Printf("round %d, player %d's turn\n", g.round, currentPlayer)

		pos, score := g.maximize(currentPlayer)
		fmt.Printf("player %d marked position %v, heuristic score： %d\n", currentPlayer, pos, score)
		g.board.b[pos.x][pos.y], currentPlayer = currentPlayer, turn(currentPlayer)
		fmt.Println("current board is:")
		g.board.Print()
	}
}

func (g *Game) maximize(currentPlayer player) (Pos, int /*heuristics score*/) {
	if winner, terminated := g.board.IsTerminated(threeInARow); terminated {
		return Pos{}, g.calcHeuristicsFromWinner(winner)
	}
	if g.round > g.simulationDepth {
		return Pos{}, g.calcHeuristicsFromBoard()
	}

	var (
		posTracker = Pos{}
		maxTracker = -math.MaxInt32
	)
	for i, row := range g.board.b {
		for j, candidate := range row {
			if candidate != none {
				continue
			}
			g.board.b[i][j] = currentPlayer
			if _, score := g.minimize(turn(currentPlayer)); score > maxTracker {
				posTracker, maxTracker = Pos{x: i, y: j}, score
			}
			g.board.b[i][j] = none
		}
	}

	return posTracker, maxTracker
}

func (g *Game) minimize(currentPlayer player) (Pos, int /*heuristics score*/) {
	if winner, terminated := g.board.IsTerminated(threeInARow); terminated {
		return Pos{}, g.calcHeuristicsFromWinner(winner)
	}
	if g.round > g.simulationDepth {
		return Pos{}, g.calcHeuristicsFromBoard()
	}

	var (
		posTracker = Pos{}
		minTracker = math.MaxInt32
	)
	for i, row := range g.board.b {
		for j, candidate := range row {
			if candidate != none {
				continue
			}
			g.board.b[i][j] = currentPlayer
			if _, score := g.maximize(turn(currentPlayer)); score < minTracker {
				posTracker, minTracker = Pos{x: i, y: j}, score
			}
			g.board.b[i][j] = none
		}
	}

	return posTracker, minTracker
}

func (g *Game) calcHeuristicsFromWinner(winner player) int {
	if winner == g.self {
		return 1
	}
	if winner == g.opponent {
		return -1
	}
	return 0
}

func (g *Game) calcHeuristicsFromBoard() int {
	panic("should not be invoked")
}

func turn(currentPlayer player) player {
	if currentPlayer == O {
		return X
	} else if currentPlayer == X {
		return O
	}
	return none
}
