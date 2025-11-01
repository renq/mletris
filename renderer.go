package main

import (
	"bytes"
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var (
	mplusFaceSource *text.GoTextFaceSource
)

func init() {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))
	if err != nil {
		log.Fatal(err)
	}
	mplusFaceSource = s
}

// Color Palette
var (
	bgColor         = color.RGBA{0x1d, 0x0f, 0x2f, 0xff} // Deep dark purple
	boardBgColor    = color.RGBA{0x2c, 0x1d, 0x40, 0xff} // Slightly lighter purple
	frameAndTextColor = color.RGBA{0xf4, 0x00, 0xff, 0xff} // Hot pink/magenta
)

type Renderer struct {
	tileSize int
	rows     int
	cols     int

	// Layout positions
	boardX     float64
	boardY     float64
	nextPieceX float64
	nextPieceY float64
	scoreX     float64
	scoreY     float64

	boardImage     *ebiten.Image
	nextPieceImage *ebiten.Image
}

func NewRenderer(tileSize, rows, cols int) *Renderer {
	boardWidth := cols * tileSize
	nextPieceWidth := 4 * tileSize
	scoreWidth := 80
	padding := 20

	totalWidth := scoreWidth + padding + boardWidth + padding + nextPieceWidth
	startX := (screenW - totalWidth) / 2

	return &Renderer{
		tileSize:       tileSize,
		rows:           rows,
		cols:           cols,
		boardImage:     ebiten.NewImage(cols*tileSize, rows*tileSize),
		nextPieceImage: ebiten.NewImage(4*tileSize, 4*tileSize),

		// Centered Layout Positions
		scoreX:     float64(startX),
		scoreY:     10,
		boardX:     float64(startX + scoreWidth + padding),
		boardY:     10,
		nextPieceX: float64(startX + scoreWidth + padding + boardWidth + padding),
		nextPieceY: 10,
	}
}

func (r *Renderer) Draw(screen *ebiten.Image, board *Board) {
	screen.Fill(bgColor)

	if board == nil {
		r.renderStartGame(screen)
		return
	}

	r.renderBoard(board, screen)
	r.renderNextPiece(board, screen)
	r.renderScore(board, screen)

	if board.paused {
		r.renderPauseOverlay(screen)
	}

	if board.gameOver {
		r.renderGameOverOverlay(screen)
	}
}

func (r *Renderer) renderBoard(board *Board, screen *ebiten.Image) {
	r.boardImage.Fill(boardBgColor)

	gridColor := adjustColor(boardBgColor, 0.8)

	// Draw vertical grid lines
	for x := 1; x < r.cols; x++ {
		vector.StrokeLine(r.boardImage, float32(x*r.tileSize), 0, float32(x*r.tileSize), float32(r.rows*r.tileSize), 1, gridColor, false)
	}
	// Draw horizontal grid lines
	for y := 1; y < r.rows; y++ {
		vector.StrokeLine(r.boardImage, 0, float32(y*r.tileSize), float32(r.cols*r.tileSize), float32(y*r.tileSize), 1, gridColor, false)
	}

	// Frame
	vector.StrokeRect(r.boardImage, 0, 0, float32(r.cols*r.tileSize), float32(r.rows*r.tileSize), 1, frameAndTextColor, true)

	// Settled tiles (flat)
	for y := 0; y < r.rows; y++ {
		for x := 0; x < r.cols; x++ {
			if board.field[y][x] != nil {
				vector.FillRect(r.boardImage, float32(x*r.tileSize), float32(y*r.tileSize), float32(r.tileSize), float32(r.tileSize), board.field[y][x], false)
			}
		}
	}

	// Current piece (flat)
	if board.currentPiece != nil {
		for _, tile := range board.currentPiece.getTiles() {
			px := float32(board.currentPiece.x*float64(r.tileSize) + float64(tile.x*r.tileSize))
			py := float32(board.currentPiece.y*float64(r.tileSize) + float64(tile.y*r.tileSize))
			vector.FillRect(r.boardImage, px, py, float32(r.tileSize), float32(r.tileSize), tile.color, false)
		}
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(r.boardX, r.boardY)
	screen.DrawImage(r.boardImage, op)
}

func (r *Renderer) renderNextPiece(b *Board, screen *ebiten.Image) {
	r.nextPieceImage.Fill(boardBgColor)
	vector.StrokeRect(r.nextPieceImage, 0, 0, float32(4*r.tileSize), float32(4*r.tileSize), 1, frameAndTextColor, true)

	if len(b.pieceQueue) > 0 && b.pieceQueue[0] != nil {
		for _, tile := range b.pieceQueue[0].getTiles() {
			// Center the piece in the 4x4 box
			px := float32((1 + tile.x) * r.tileSize)
			py := float32((1 + tile.y) * r.tileSize)
			vector.FillRect(r.nextPieceImage, px, py, float32(r.tileSize), float32(r.tileSize), tile.color, false)
		}
	}

	opNext := &ebiten.DrawImageOptions{}
	opNext.GeoM.Translate(r.nextPieceX, r.nextPieceY)
	screen.DrawImage(r.nextPieceImage, opNext)
}

func (r *Renderer) renderScore(b *Board, screen *ebiten.Image) {
	// --- Score ---
	scoreTitleOp := &text.DrawOptions{}
	scoreTitleOp.GeoM.Translate(r.scoreX, r.scoreY)
	scoreTitleOp.ColorScale.ScaleWithColor(frameAndTextColor)
	text.Draw(screen, "SCORE", &text.GoTextFace{Source: mplusFaceSource, Size: 12}, scoreTitleOp)

	scoreValueOp := &text.DrawOptions{}
	scoreValueOp.GeoM.Translate(r.scoreX, r.scoreY+15)
	scoreValueOp.ColorScale.ScaleWithColor(frameAndTextColor)
	scoreStr := fmt.Sprintf("%d", b.Score)
	text.Draw(screen, scoreStr, &text.GoTextFace{Source: mplusFaceSource, Size: 12}, scoreValueOp)

	// --- Level ---
	levelTitleOp := &text.DrawOptions{}
	levelTitleOp.GeoM.Translate(r.scoreX, r.scoreY+45)
	levelTitleOp.ColorScale.ScaleWithColor(frameAndTextColor)
	text.Draw(screen, "LEVEL", &text.GoTextFace{Source: mplusFaceSource, Size: 12}, levelTitleOp)

	levelValueOp := &text.DrawOptions{}
	levelValueOp.GeoM.Translate(r.scoreX, r.scoreY+60)
	levelValueOp.ColorScale.ScaleWithColor(frameAndTextColor)
	levelStr := fmt.Sprintf("%d", b.Level)
	text.Draw(screen, levelStr, &text.GoTextFace{Source: mplusFaceSource, Size: 12}, levelValueOp)
}

func (r *Renderer) renderStartGame(screen *ebiten.Image) {
	textString := "Press [Enter] to Start"
	op := &text.DrawOptions{}
	op.GeoM.Translate(float64(screenW)/2-100, float64(screenH)/2)
	op.ColorScale.ScaleWithColor(frameAndTextColor)
	text.Draw(screen, textString, &text.GoTextFace{
		Source: mplusFaceSource,
		Size:   20,
	}, op)
}

func (r *Renderer) renderPauseOverlay(screen *ebiten.Image) {
	textString := "Paused"
	op := &text.DrawOptions{}
	op.GeoM.Translate(float64(screenW)/2-40, float64(screenH)/2)
	op.ColorScale.ScaleWithColor(frameAndTextColor)
	text.Draw(screen, textString, &text.GoTextFace{
		Source: mplusFaceSource,
		Size:   24,
	}, op)
}

func (r *Renderer) renderGameOverOverlay(screen *ebiten.Image) {
	textString := "GAME OVER"
	op := &text.DrawOptions{}
	op.GeoM.Translate(float64(screenW)/2-60, float64(screenH)/2-30)
	op.ColorScale.ScaleWithColor(frameAndTextColor)
	text.Draw(screen, textString, &text.GoTextFace{
		Source: mplusFaceSource,
		Size:   24,
	}, op)

	op.GeoM.Translate(0, 30)
	text.Draw(screen, "Press [Enter] to start again", &text.GoTextFace{
		Source: mplusFaceSource,
		Size:   12,
	}, op)
}

// adjustColor is a helper to create a lighter or darker version of a color.
func adjustColor(c color.Color, factor float32) color.Color {
	r, g, b, a := c.RGBA()
	return color.RGBA{
		R: uint8(min(255, float32(r>>8)*factor)),
		G: uint8(min(255, float32(g>>8)*factor)),
		B: uint8(min(255, float32(b>>8)*factor)),
		A: uint8(a >> 8),
	}
}