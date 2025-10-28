package main

import (
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
	data [][]Tile
}

type FallingPiece struct{
	piece Piece
	state int
	x float64
	y float64
}

func (fp FallingPiece) rotate() FallingPiece {
	fp.state = (fp.state + 1) % len(fp.piece.data)

	return fp
}

func (fp *FallingPiece) getTiles() []Tile {
	return fp.piece.data[fp.state]
}


type Board struct{
	backgroundColor color.Color // consider moving it away from this struct
	paused bool
	gameOver bool
	tickNumber int
	ticksPerSecond int
	frameColor color.Color
	field Field
	currentPiece *FallingPiece
	pieceQueue []*FallingPiece
	tiles []Piece
}

func NewBoard(rows int, cols int, ticksPerSecond int) *Board {
	b := &Board{
		backgroundColor: color.RGBA{0x3e, 0x22, 0x2d, 0xff}, // TODO move it from here
		frameColor: color.RGBA{0xbb, 0xad, 0xa0, 0xff}, // TODO move it from here

		ticksPerSecond: ticksPerSecond,
		field: createField(rows, cols),
		tiles: buildTiles(),
	}

	b.pieceQueue = []*FallingPiece{
		b.generateRandomPiece(),
		b.generateRandomPiece(),
		b.generateRandomPiece(),
	}

	b.currentPiece = b.newPiece()

	return b
}

func (b *Board) TogglePause() {
	b.paused = !b.paused
}

func (b *Board) Tick() {
	if b.isStopped() {
		return
	}

	b.tickNumber++
	
	if b.tickNumber >= b.ticksPerSecond {
		b.MoveDown()
    }
}

func (b *Board) MoveRight() {
	if b.isStopped() {
		return
	}

	if b.canMove(b.currentPiece, 1) {
		b.currentPiece.x += 1.0
	}
}

func (b *Board) MoveLeft() {
	if b.isStopped() {
		return
	}

	// TODO There is a bug here when moving a block under another block. I don't know how to reproduce it yet.
	if b.canMove(b.currentPiece, -1) {
		b.currentPiece.x -= 1.0
	}
}

func (b *Board) MoveDown() {
	if b.isStopped() {
		return
	}

	b.tickNumber = 0
	if (b.collisionDetected(b.currentPiece)) {
		b.addCurrentPieceToTheBoard()
		b.currentPiece = b.newPiece()
	}

	b.currentPiece.y += 1.0
}

func (b *Board) Fall() {
	if b.isStopped() {
		return
	}

	for !b.collisionDetected(b.currentPiece) {
		b.currentPiece.y += 1.0
	}

	b.addCurrentPieceToTheBoard()
	b.currentPiece = b.newPiece()
	b.tickNumber = 0
}

func (b *Board) Rotate() {
	if b.isStopped() {
		return
	}

	// TODO check for collisions
	rotated := b.currentPiece.rotate()

	minX := cols - 1
	maxX := 0

	// wall kick check
	for _, tile := range rotated.getTiles() {
		minX = min(minX, int(rotated.x) + tile.x)
		maxX = max(maxX, int(rotated.x) + tile.x)
	}

	if minX < 0 {
		rotated.x += float64(-minX)
	}

	if maxX >= cols {
		rotated.x -= float64(maxX - (cols - 1))
	}	

	// final collision check
	if (b.collisionDetected(&rotated)) {
		return
	}

	b.currentPiece = &rotated
}

func (b *Board) isStopped() bool {
	return b.paused != false || b.gameOver == true
}

func (b *Board) newPiece() *FallingPiece {
	piece := b.pieceQueue[0]

	b.pieceQueue = b.pieceQueue[1:]
	b.pieceQueue = append(b.pieceQueue, b.generateRandomPiece())

	if (b.collisionDetected(piece)) {
		b.gameOver = true
	}

	return piece
}

func (b *Board) generateRandomPiece() *FallingPiece {
	id := rand.Intn(len(b.tiles))
	piece := &FallingPiece{
		piece: b.tiles[id],
		x: 4.,
		y: 1.,
	}
	
	return piece
}

// TODO try to merge canMove and collisionDetected. They are very similar.
func (b *Board) canMove(piece *FallingPiece, direction int) bool {
	for _, tile := range piece.getTiles() {
		newX := int(piece.x) + tile.x + direction

		if newX < 0 || newX >= cols {
			return false
		}

		if (b.field[int(piece.y)][newX] != nil) {
			return false
		}
	}

	return true
}

func (b *Board) collisionDetected(piece *FallingPiece) bool {
	for _, tile := range piece.getTiles() {
		newY := int(piece.y) + tile.y + 1

		if newY >= rows {
			return true
		}

		if (b.field[newY][int(piece.x) + tile.x] != nil) {
			return true
		}
	}

	return false
}

func (b *Board) addCurrentPieceToTheBoard() {
	// Add to board
	for _, tile := range b.currentPiece.getTiles() {
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
