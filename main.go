package main

import (
	"bytes"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"

	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
)

const (
	ticksPerSecond = 60
	pressDelayTicks = ticksPerSecond / 6
	pressRepeatIntervalTicks = ticksPerSecond / 30
	screenW = 320
	screenH = 240
	rows = 24
	cols = 10
	tileSize = 9
)

type Game struct{
	backgroundColor color.Color
	board *Board
}

var (
	mplusFaceSource *text.GoTextFaceSource
	boardImage = ebiten.NewImage(cols * tileSize, rows * tileSize)
	nextPieceImage = ebiten.NewImage(4 * tileSize, 4 * tileSize)
	tileImage  = ebiten.NewImage(tileSize - 1, tileSize - 1)
)

func NewGame() *Game {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))
	if err != nil {
		log.Fatal(err)
	}
	mplusFaceSource = s

	return &Game{
		backgroundColor: color.RGBA{0x3e, 0x12, 0x2d, 0xff},
	}
}

func keyPressAndMove(key ebiten.Key) bool {
	return inpututil.IsKeyJustPressed(key) || 
		inpututil.KeyPressDuration(key) > pressDelayTicks && 
		inpututil.KeyPressDuration(key) % pressRepeatIntervalTicks == 0
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		g.board = NewBoard(rows, cols, ticksPerSecond)
	}

	if g.board != nil {
		b := g.board


		b.Tick()
		
		if keyPressAndMove(ebiten.KeyArrowLeft) || keyPressAndMove(ebiten.KeyA) {
			b.MoveLeft()
		}

		if keyPressAndMove(ebiten.KeyArrowRight) || keyPressAndMove(ebiten.KeyD) {
			b.MoveRight()
		}

		if keyPressAndMove(ebiten.KeyArrowDown) || keyPressAndMove(ebiten.KeyS) {
			b.MoveDown()
		}

		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			b.Fall()
		}

		if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) || inpututil.IsKeyJustPressed(ebiten.KeyW) {
			b.Rotate()
		}

		if inpututil.IsKeyJustPressed(ebiten.KeyP) {
			b.TogglePause()
		}
	}

	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	// TODO refactor this mess! It's too big, too hard to read.
	screen.Fill(g.backgroundColor)

	if g.board == nil {
		// The text you want to draw
		textString := "Press [Enter] to Start"

		// Options for drawing (position, color, etc.)
		op := &text.DrawOptions{}
		op.GeoM.Translate(50, 50) // X=50, Y=50 position (top-left of text bounds, roughly)
		op.ColorScale.ScaleWithColor(color.White)

		// Draw the text using the source and size defined earlier
		text.Draw(screen, textString, &text.GoTextFace{
			Source: mplusFaceSource,
			Size:   20,
		}, op)
	}

	if g.board != nil {
		// Board drawing - start	
		// Frame
		board := g.board

		boardImage.Fill(board.backgroundColor)
		vector.StrokeLine(boardImage, 0, 0, 0, tileSize * rows, 1, board.frameColor, true)
		vector.StrokeLine(boardImage, 0, 0, tileSize * cols, 0, 1, board.frameColor, true)
		vector.StrokeLine(boardImage, tileSize * cols, 0, tileSize * cols, tileSize * rows, 1, board.frameColor, true)
		vector.StrokeLine(boardImage, 0, tileSize * rows, tileSize * cols, tileSize * rows, 1, board.frameColor, true)

		// Tiles
		for y := 0; y < rows; y++ {
			for x := 0; x < cols; x++ {
				op := &ebiten.DrawImageOptions{}
				if (board.field[y][x] == nil) {
					continue
				}
				tileImage.Fill(board.field[y][x])
				
				op.GeoM.Translate(float64(x*tileSize) + tileSize/5, float64(y*tileSize) + tileSize/5)
				boardImage.DrawImage(tileImage, op)
			}
		}

		// Current piece
		for _, tile := range board.currentPiece.getTiles() {
			op := &ebiten.DrawImageOptions{}
			tileImage.Fill(tile.color)
			op.GeoM.Translate(
				board.currentPiece.x * tileSize + float64(tile.x * tileSize) + tileSize/5, 
				board.currentPiece.y * tileSize + float64(tile.y * tileSize) + tileSize/5,
			)
			boardImage.DrawImage(tileImage, op)
		}

		// Board drawing - end

		
		// Draw board image to screen
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(10, 10)
		screen.DrawImage(boardImage, op)

		// Pause overlay
		if (board.paused) {
			textString := "Paused"

			op := &text.DrawOptions{}
			op.GeoM.Translate(100, 100)
			op.ColorScale.ScaleWithColor(color.White)

			text.Draw(screen, textString, &text.GoTextFace{
				Source: mplusFaceSource,
				Size:   24,
			}, op)
		}

		// Game over overlay
		if (board.gameOver) {
			textString := "Game over"

			op := &text.DrawOptions{}
			op.GeoM.Translate(100, 100)
			op.ColorScale.ScaleWithColor(color.White)

			text.Draw(screen, textString, &text.GoTextFace{
				Source: mplusFaceSource,
				Size:   24,
			}, op)


			op.GeoM.Translate(0, 30)
			text.Draw(screen, "[enter] to start a new game", &text.GoTextFace{
				Source: mplusFaceSource,
				Size:   12,
			}, op)
		}

		// Next piece preview
		nextPieceImage.Fill(color.RGBA{0x20, 0x20, 0x20, 0xff})
		
		for _, tile := range board.pieceQueue[0].getTiles() {
			op := &ebiten.DrawImageOptions{}
			tileImage.Fill(tile.color)
			op.GeoM.Translate(
				float64((1 + tile.x) * tileSize) + tileSize/5, 
				float64((1 + tile.y) * tileSize) + tileSize/5,
			)
			nextPieceImage.DrawImage(tileImage, op)
		}

		opNext := &ebiten.DrawImageOptions{}
		opNext.GeoM.Translate(float64(screenW - (4 * tileSize) - 10), 10)
		screen.DrawImage(nextPieceImage, opNext)
	}
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return screenW, screenH
}

func main() {
	game := NewGame()

    ebiten.SetWindowSize(1024, 768)
    ebiten.SetWindowTitle("Mletris")

    if err := ebiten.RunGame(game); err != nil {
        log.Fatal(err)
    }
}
