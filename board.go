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


type Board struct {
	paused         bool
	gameOver       bool
	tickNumber     int
	Score          int
	Level          int
	totalNumberOfLinesCleared int
	linesCleared   int
	field          Field
	currentPiece   *FallingPiece
	pieceQueue     []*FallingPiece
	tiles          []Piece
}

func NewBoard(rows int, cols int) *Board {
	b := &Board{
		Level:          0,
		linesCleared:   0,
		field:          createField(rows, cols),
		tiles:          buildTiles(),
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

	b.nextLevelIfNeeded()

	if b.timeToDrop() {
		b.MoveDown()
	}
}

func (b *Board) MoveRight() {
	if b.isStopped() {
		return
	}

	if !b.checkCollision(b.currentPiece, 1, 0) {
		b.currentPiece.x += 1.0
	}
}

func (b *Board) MoveLeft() {
	if b.isStopped() {
		return
	}

	// TODO There is a bug here when moving a block under another block. I don't know how to reproduce it yet.
	if !b.checkCollision(b.currentPiece, -1, 0) {
		b.currentPiece.x -= 1.0
	}
}

func (b *Board) MoveDown() {
	if b.isStopped() {
		return
	}

	b.tickNumber = 0
	if b.checkCollision(b.currentPiece, 0, 1) {
		b.addCurrentPieceToTheBoard()
		b.currentPiece = b.newPiece()
		return
	}

	b.currentPiece.y += 1.0
}

func (b *Board) Fall() {
	if b.isStopped() {
		return
	}

	for !b.checkCollision(b.currentPiece, 0, 1) {
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
	if b.checkCollision(&rotated, 0, 0) {
		return
	}

	b.currentPiece = &rotated
}

func (b *Board) isStopped() bool {
	return b.paused || b.gameOver
}

func (b *Board) newPiece() *FallingPiece {
	piece := b.pieceQueue[0]

	b.pieceQueue = b.pieceQueue[1:]
	b.pieceQueue = append(b.pieceQueue, b.generateRandomPiece())

	if b.checkCollision(piece, 0, 0) {
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

func (b *Board) checkCollision(p *FallingPiece, xOffset, yOffset float64) bool {
    for _, tile := range p.getTiles() {
        newX := int(p.x+xOffset) + tile.x
        newY := int(p.y+yOffset) + tile.y

        if newX < 0 || newX >= cols || newY >= rows {
            return true
        }

        if newY >= 0 && b.field[newY][newX] != nil {
            return true
        }
    }
    return false
}

func (b *Board) addCurrentPieceToTheBoard() {
	// Add to board
	for _, tile := range b.currentPiece.getTiles() {
		newY := int(b.currentPiece.y) + tile.y
		b.field[newY][int(b.currentPiece.x)+tile.x] = tile.color
	}

	// Clean full lines and count them
	clearedCount := 0
	writeRow := rows - 1
	for readRow := rows - 1; readRow >= 0; readRow-- {
		isFull := true
		for x := 0; x < cols; x++ {
			if b.field[readRow][x] == nil {
				isFull = false
				break
			}
		}

		if (isFull) {
			clearedCount++
		} else {
			if readRow != writeRow {
				b.field[writeRow] = b.field[readRow]
			}
			writeRow--
		}
	}

	if clearedCount > 0 {
		b.totalNumberOfLinesCleared += clearedCount
		b.addScore(clearedCount)
		b.linesCleared += clearedCount
		// Level up every 10 lines
		if b.linesCleared >= 10 {
			b.Level++
			b.linesCleared -= 10
		}
	}

	// Fill the cleared lines at the top with new empty rows
	for y := writeRow; y >= 0; y-- {
		b.field[y] = make([]color.Color, cols)
	}
}

func (b *Board) timeToDrop() bool {
	level := b.Level
	if level >= len(framePerDrop) {
		level = len(framePerDrop) - 1
	}
	return b.tickNumber >= framePerDrop[level]
}

func (b *Board) addScore(lines int) {
	// Standard Tetris scoring
	baseScores := []int{0, 40, 100, 300, 1200}
	scoreIndex := lines
	if scoreIndex > 4 {
		scoreIndex = 4
	}

	b.Score += baseScores[scoreIndex] * (b.Level + 1)
}

func (b *Board) nextLevelIfNeeded() {
	for b.totalNumberOfLinesCleared >= (b.Level+1)*10 {
		b.Level++
	}
}
