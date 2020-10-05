package game

import (
	"fmt"
)

type player int

const (
	none player = iota
	O
	X
)

type Pos struct {
	x, y int
}

var positionHolder = Pos{x: -1, y: -1}

type Board struct {
	b [][]player
}

func NewBoard(dimension int) *Board {
	b := make([][]player, dimension)
	for i := range b {
		b[i] = make([]player, dimension)
	}
	return &Board{b: b}
}

func (b *Board) IsTerminated(factor int) (player, bool) {
	var hasUnfilledSlot bool
	for i, row := range b.b {
		for j, candidate := range row {
			if candidate == none {
				hasUnfilledSlot = true
				continue
			}
		directions:
			for _, delta := range []Pos{{0, 1}, {1, 0}, {1, 1}, {1, -1}} {
				for k := 1; k < factor; k++ {
					next := Pos{x: i + k*delta.x, y: j + k*delta.y}
					if !b.IsValidPosition(next) {
						continue directions
					}
					if b.b[next.x][next.y] != candidate {
						continue directions
					}
				}
				return candidate, true
			}
		}
	}
	return none, !hasUnfilledSlot
}

func (b *Board) calcPointsFor(target player) int {
	var points int
	for i, row := range b.b {
		for j, candidate := range row {
			if candidate != target {
				continue
			}

			for _, delta := range []Pos{{0, 1}, {1, 0}, {1, 1}, {1, -1}} {
				next := Pos{x: i + delta.x, y: j + delta.y}
				if !b.IsValidPosition(next) {
					continue
				}
				if b.b[next.x][next.y] == target {
					points += 1
				}
			}
		}
	}
	return points
}

func (b *Board) IsValidPosition(pos Pos) bool {
	return pos.x >= 0 && pos.x < len(b.b) && pos.y >= 0 && pos.y < len(b.b[pos.x])
}

func (b *Board) Print() {
	for _, row := range b.b {
		fmt.Println(row)
	}
}
