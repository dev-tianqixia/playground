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
	round           int
	simulationDepth int // unused
}

func NewGame(board *Board, simulationDepth int) *Game {
	return &Game{
		board:           board,
		simulationDepth: simulationDepth,
	}
}

// trigger the simulation
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

		pos, score := g.maximize(currentPlayer, 0, -math.MaxInt32-1, math.MaxInt32+1)
		fmt.Printf("player %d marked position %v, heuristic score： %d\n", currentPlayer, pos, score)
		g.board.b[pos.x][pos.y], currentPlayer = currentPlayer, turn(currentPlayer)
		fmt.Println("current board is:")
		g.board.Print()
		fmt.Println("")
	}
}

func (g *Game) maximize(currentPlayer player, depth, min, max int) (Pos, int /*heuristics score*/) {
	if winner, terminated := g.board.IsTerminated(threeInARow); terminated {
		return positionHolder, g.calcHeuristicsFromWinner(winner)
	}
	if depth > g.simulationDepth {
		return positionHolder, g.calcHeuristicsFromBoard()
	}

	var (
		posTracker = positionHolder
		maxTracker = min
	)
	for i, row := range g.board.b {
		for j, candidate := range row {
			if candidate != none {
				continue
			}

			g.board.b[i][j] = currentPlayer
			_, score := g.minimize(turn(currentPlayer), depth+1, maxTracker, max)
			//if depth == 0 && g.self == X {
			//	fmt.Printf("position: (%d, %d), score: %d\n", i, j, score)
			//}
			if score > max {
				g.board.b[i][j] = none
				return positionHolder, max
			} else if score > maxTracker {
				posTracker, maxTracker = Pos{x: i, y: j}, score
			}
			g.board.b[i][j] = none
		}
	}

	return posTracker, maxTracker
}

func (g *Game) minimize(currentPlayer player, depth, min, max int) (Pos, int /*heuristics score*/) {
	if winner, terminated := g.board.IsTerminated(threeInARow); terminated {
		return positionHolder, g.calcHeuristicsFromWinner(winner)
	}
	if depth > g.simulationDepth {
		return positionHolder, g.calcHeuristicsFromBoard()
	}

	var (
		posTracker = positionHolder
		minTracker = max
	)
	for i, row := range g.board.b {
		for j, candidate := range row {
			if candidate != none {
				continue
			}

			g.board.b[i][j] = currentPlayer
			if _, score := g.maximize(turn(currentPlayer), depth+1, min, minTracker); score < min {
				g.board.b[i][j] = none
				return positionHolder, min
			} else if score < minTracker {
				posTracker, minTracker = Pos{x: i, y: j}, score
			}
			g.board.b[i][j] = none
		}
	}

	return posTracker, minTracker
}

// heuristics are calculated from g.self's perspective
func (g *Game) calcHeuristicsFromWinner(winner player) int {
	if winner == g.self {
		return math.MaxInt32
	}
	if winner == g.opponent {
		return -math.MaxInt32
	}
	return 0
}

// heuristics are calculated from g.self's perspective
func (g *Game) calcHeuristicsFromBoard() int {
	return g.board.calcPointsFor(g.self)
}

func turn(currentPlayer player) player {
	if currentPlayer == O {
		return X
	} else if currentPlayer == X {
		return O
	}
	return none
}
