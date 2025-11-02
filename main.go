package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	screenW        = 320
	screenH        = 240
	rows           = 24
	cols           = 10
	tileSize       = 9
)

type Game struct {
	board        *Board
	inputHandler *InputHandler
	renderer     *Renderer
}

func NewGame() *Game {
	g := &Game{
		inputHandler: &InputHandler{},
		renderer:     NewRenderer(tileSize, rows, cols),
	}
	return g
}

func (g *Game) Update() error {
	// Global input handling (creating a new game)
	if (g.board == nil || g.board.gameOver) && inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		g.board = NewBoard(rows, cols)
	}

	// Delegate board-related input to the handler
	g.inputHandler.Update(g.board)

	// Update game state
	if g.board != nil {
		g.board.Tick()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Delegate all drawing to the renderer
	g.renderer.Draw(screen, g.board)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenW, screenH
}

func main() {
	ebiten.SetWindowSize(1024, 768)
	ebiten.SetWindowTitle("Mletris")

	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}