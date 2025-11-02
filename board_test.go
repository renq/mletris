package main

import (
	"image/color"
	"strings"
	"testing"
)

func TestRotate(t *testing.T) {
	b := NewBoard(rows, cols)

	// Manually set the current piece to a T piece to avoid randomness
	b.currentPiece = &FallingPiece{
		piece: b.tiles[2], // T piece
		x:     4.,
		y:     1.,
		state: 0,
	}

	initialState := b.currentPiece.state
	b.Rotate()
	finalState := b.currentPiece.state

	if finalState == initialState {
		t.Errorf("Expected piece to rotate, but it did not. Initial state: %d, Final state: %d", initialState, finalState)
	}

	expectedState := (initialState + 1) % len(b.currentPiece.piece.data)
	if finalState != expectedState {
		t.Errorf("Expected state to be %d after rotation, but got %d", expectedState, finalState)
	}
}

func TestRotate_NoSpace(t *testing.T) {
	b := NewBoard(rows, cols)

	// Use the helper to create a board with limited space
	layout := `
..........
..........
..........
..........
xxxx...xxx
xxxx...xxx
xxxx...xxx
xxxx...xxx
xxxx...xxx
xxxx...xxx
xxxx...xxx
xxxx...xxx
xxxx...xxx
xxxx...xxx
xxxx...xxx
xxxx...xxx
xxxx...xxx
xxxx...xxx
xxxx...xxx
xxxx...xxx
xxxx...xxx
xxxx...xxx
xxxx...xxx
xxxx...xxx`
	fillBoardFromString(b, layout)

	// Manually set the current piece to an I piece in the gap
	b.currentPiece = &FallingPiece{
		piece: b.tiles[0], // I piece
		x:     5.,
		y:     10.,
		state: 0, // Vertical state
	}

	initialState := b.currentPiece.state
	b.Rotate()
	finalState := b.currentPiece.state

	if finalState != initialState {
		t.Errorf("Expected piece not to rotate, but it did. Initial state: %d, Final state: %d", initialState, finalState)
	}
}

// fillBoardFromString populates the board's field based on a string representation.
// 'x' represents a block, '.' represents an empty space.
func fillBoardFromString(board *Board, layout string) {
	lines := strings.Split(strings.TrimSpace(layout), "\n")

	for r, line := range lines {
		if r < rows {
			for c, char := range line {
				if c < cols {
					if char == 'x' {
						board.field[r][c] = color.White
					}
				}
			}
		}
	}
}
