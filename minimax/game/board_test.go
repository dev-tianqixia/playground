package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBoard_IsTerminated(t *testing.T) {
	// empty board should be non-terminated
	board := NewBoard(3)
	winner, terminated := board.IsTerminated(3)
	assert.Equal(t, false, terminated)
	assert.Equal(t, none, winner)
	// basic unterminated case 1
	board = &Board{b: [][]player{
		{1, 2, 0},
		{0, 1, 1},
		{2, 0, 2},
	}}
	winner, terminated = board.IsTerminated(3)
	assert.Equal(t, false, terminated)
	assert.Equal(t, none, winner)
	// basic terminated case 1
	board = &Board{b: [][]player{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	}}
	winner, terminated = board.IsTerminated(3)
	assert.Equal(t, true, terminated)
	assert.Equal(t, O, winner)
	// basic terminated case 2
	board = &Board{b: [][]player{
		{1, 0, 1},
		{2, 2, 2},
		{0, 1, 1},
	}}
	winner, terminated = board.IsTerminated(3)
	assert.Equal(t, true, terminated)
	assert.Equal(t, X, winner)
	// basic terminated case 3
	board = &Board{b: [][]player{
		{1, 2, 1},
		{2, 1, 2},
		{2, 1, 2},
	}}
	winner, terminated = board.IsTerminated(3)
	assert.Equal(t, true, terminated)
	assert.Equal(t, none, winner)
}
