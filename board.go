package main

import (
	"fmt"
	"image/color"
	"math/rand"
)

type Field [][]color.Color

func createField(rows int, cols int) Field{
	matrix := make(Field, rows)
	for i := range matrix {
		matrix[i] = make([]color.Color, cols)
	}

	return matrix
}

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
	backgroundColor color.Color // consider moving it away from this struct
	tickNumber int
	ticksPerSecond int
	frameColor color.Color
	field Field
	currentPiece *FallingPiece
	tiles [2]Piece
}

func NewBoard(rows int, cols int, ticksPerSecond int) *Board {
	b := &Board{
		backgroundColor: color.RGBA{0x3e, 0x22, 0x2d, 0xff}, // TODO move it from here
		frameColor: color.RGBA{0xbb, 0xad, 0xa0, 0xff}, // TODO move it from here

		ticksPerSecond: ticksPerSecond,
		field: createField(rows, cols),
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
	
	if b.tickNumber >= b.ticksPerSecond {
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
	if (b.collisionDetected()) {
		b.addToBoard()
		b.currentPiece = b.newPiece()

		return true
	}

	b.currentPiece.y += 1.0

	return false
}

func (b *Board) Fall() {
	fmt.Printf("Move down\n")
	for !b.collisionDetected() {
		b.currentPiece.y += 1.0
	}
	fmt.Printf("Moved down y = %d\n", int(b.currentPiece.y))

	b.addToBoard()
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

		if newX < 0 || newX >= cols {
			return false
		}

		if (b.field[int(b.currentPiece.y)][newX] != nil) {
			return false
		}
	}

	return true
}

func (b *Board) collisionDetected() bool {
	for _, tile := range b.currentPiece.piece.tiles {
		newY := int(b.currentPiece.y) + tile.y + 1

		if newY >= rows {
			return true
		}

		if (b.field[newY][int(b.currentPiece.x) + tile.x] != nil) {
			return true
		}
	}

	return false
}

func (b *Board) addToBoard() {
	// Add to board
	for _, tile := range b.currentPiece.piece.tiles {
		newY := int(b.currentPiece.y) + tile.y

		b.field[newY][int(b.currentPiece.x) + tile.x] = tile.color
	}

	// Clean full lines
	for y := 0; y < rows; y++ {
		skip := false

		for x := 0; x < cols; x++ {
			if (b.field[y][x] == nil) {
				skip = true
				break
			}
		}

		if skip {
			continue
		}

		b.field = append(b.field[:y], b.field[y+1:]...)
		b.field = append([][]color.Color{make([]color.Color, cols)}, b.field...)
	}
}
