package main

import (
	"log"

	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	ticksPerSecond = 60
	screenWidth = 640
	screenHeight = 480
	ROWS = 24
	COLS = 10
	TILE_SIZE = 9
)




type Game struct{
	backgroundColor color.Color
	board *Board
}

func NewGame() *Game {
	return &Game{
		backgroundColor: color.RGBA{0x3e, 0x12, 0x2d, 0xff},
		board: NewBoard(),
	}
}

var (
	boardImage      = ebiten.NewImage(COLS * TILE_SIZE, ROWS * TILE_SIZE)
	tileImage       = ebiten.NewImage(TILE_SIZE - 1, TILE_SIZE - 1)
)

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	b := g.board

	b.Tick()

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {
		b.MoveRight()
	}
	
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {
		b.MoveLeft()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		b.Fall()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		b.MoveDown()
	}

	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(g.backgroundColor)

	// Board drawing - start
	// Frame
	board := g.board

	boardImage.Fill(board.backgroundColor)
	vector.StrokeLine(boardImage, 0, 0, 0, TILE_SIZE * ROWS, 1, board.frameColor, true)
	vector.StrokeLine(boardImage, 0, 0, TILE_SIZE * COLS, 0, 1, board.frameColor, true)
	vector.StrokeLine(boardImage, TILE_SIZE * COLS, 0, TILE_SIZE * COLS, TILE_SIZE * ROWS, 1, board.frameColor, true)
	vector.StrokeLine(boardImage, 0, TILE_SIZE * ROWS, TILE_SIZE * COLS, TILE_SIZE * ROWS, 1, board.frameColor, true)

	// Tiles
	for y := 0; y < ROWS; y++ {
		for x := 0; x < COLS; x++ {
			op := &ebiten.DrawImageOptions{}
			if (board.field[y][x] == nil) {
				continue
			}
			tileImage.Fill(board.field[y][x])
			
			op.GeoM.Translate(float64(x*TILE_SIZE) + TILE_SIZE/5, float64(y*TILE_SIZE) + TILE_SIZE/5)
			boardImage.DrawImage(tileImage, op)
		}
	}

	// Current piece
	for _, tile := range board.currentPiece.piece.tiles {
		op := &ebiten.DrawImageOptions{}
		tileImage.Fill(tile.color)
		op.GeoM.Translate(
			board.currentPiece.x * TILE_SIZE + float64(tile.x * TILE_SIZE) + TILE_SIZE/5, 
			board.currentPiece.y * TILE_SIZE + float64(tile.y * TILE_SIZE) + TILE_SIZE/5,
		)
		boardImage.DrawImage(tileImage, op)
	}
	// Board drawing - end

	// Draw board image to screen
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(10, 10)
	screen.DrawImage(boardImage, op)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	game := NewGame()

    // Specify the window size as you like. Here, a doubled size is specified.
    ebiten.SetWindowSize(800, 600)
    ebiten.SetWindowTitle("Mletris")
    // Call ebiten.RunGame to start your game loop.
    if err := ebiten.RunGame(game); err != nil {
        log.Fatal(err)
    }
}
