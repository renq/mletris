package main

import (
	"fmt"
	"image/color"
	"math/rand"
)

type Field [ROWS][COLS]color.Color

type Tile struct{
	x int
	y int
	color color.Color
}
	
type Piece struct{
	tiles []Tile
}

type FallingPiece struct{
	piece Piece
	x float64
	y float64
}



type Board struct{
	backgroundColor color.Color
	tickNumber int
	frameColor color.Color
	field *Field
	currentPiece *FallingPiece
	tiles [2]Piece
}

func NewBoard() *Board {
	b := &Board{
		backgroundColor: color.RGBA{0x3e, 0x22, 0x2d, 0xff},
		frameColor: color.RGBA{0xbb, 0xad, 0xa0, 0xff},
		field: &Field{},
		tiles: [2]Piece{
			Piece{
				tiles: []Tile{
					Tile{x:0, y:0, color: color.RGBA{0x22, 0xff, 0x4b, 0xff}},
				},
			},
			Piece{
				tiles: []Tile{
					Tile{x:0, y:0, color: color.RGBA{0xff, 0x22, 0x4b, 0xff}},
					Tile{x:1, y:0, color: color.RGBA{0xff, 0x22, 0x4b, 0xff}},
				},
			},
		},
	}

	b.currentPiece = b.newPiece()

	return b
}

func (b *Board) Tick() {
	b.tickNumber++
	
	if b.tickNumber >= ticksPerSecond {
		b.MoveDown()
    }
}

func (b *Board) MoveRight() bool {
	if b.canMove(1) {
		b.currentPiece.x += 1.0
		return true
	}
	
	return false
}

func (b *Board) MoveLeft() bool {
	if b.canMove(-1) {
		b.currentPiece.x -= 1.0
		return true
	}
	
	return false
}

func (b *Board) MoveDown() bool {
	b.tickNumber = 0
	if (b.collisionDetected(b.currentPiece)) {
		b.addToBoard(b.currentPiece)
		b.currentPiece = b.newPiece()

		return true
	}

	b.currentPiece.y += 1.0

	return false
}

func (b *Board) Fall() {
	fmt.Printf("Move down\n")
	for !b.collisionDetected(b.currentPiece) {
		b.currentPiece.y += 1.0
	}
	fmt.Printf("Moved down y = %d\n", int(b.currentPiece.y))

	b.addToBoard(b.currentPiece)
	b.currentPiece = b.newPiece()
	b.tickNumber = 0
}

func (b *Board) newPiece() *FallingPiece {
	id := rand.Intn(len(b.tiles))
	return &FallingPiece{
		piece: b.tiles[id],
		x: 4.,
		y: 1.,
	}
}

func (b *Board) canMove(direction int) bool {
	for _, tile := range b.currentPiece.piece.tiles {
		newX := int(b.currentPiece.x) + tile.x + direction

		if newX < 0 || newX >= COLS {
			return false
		}

		if (b.field[int(b.currentPiece.y)][newX] != nil) {
			return false
		}
	}

	return true
}

func (b *Board) collisionDetected(piece *FallingPiece) bool {
	for _, tile := range piece.piece.tiles {
		newY := int(piece.y) + tile.y + 1

		if newY >= ROWS {
			return true
		}

		if (b.field[newY][int(piece.x) + tile.x] != nil) {
			return true
		}
	}

	return false
}

func (b *Board) addToBoard(piece *FallingPiece) {
	// Add to board
	for _, tile := range piece.piece.tiles {
		newY := int(piece.y) + tile.y

		b.field[newY][int(piece.x) + tile.x] = tile.color
	}

	// Try to clean lines
	for y := 0; y < ROWS; y++ {
		skip := false
		fmt.Printf("Check line y = %d\n", y)
		for x := 0; x < COLS; x++ {
			if (b.field[y][x] == nil) {
				skip = true
				break
			}
		}

		if skip {
			continue
		}
		
	    fmt.Printf("Clear line y = %d\n", y)

		// Clear line
		for ty := y; ty > 0; ty-- {
			for x := 0; x < COLS; x++ {
				b.field[ty][x] = b.field[ty-1][x]
			}
		}
		// Clear top line
		for x := 0; x < COLS; x++ {
			b.field[0][x] = nil
		}
	}
}
