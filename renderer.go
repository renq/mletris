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

type Renderer struct {
	tileSize int
	rows     int
	cols     int

	backgroundColor      color.Color
	boardBackgroundColor color.Color
	boardFrameColor      color.Color

	boardImage     *ebiten.Image
	nextPieceImage *ebiten.Image
	scoreImage     *ebiten.Image
	tileImage      *ebiten.Image
}

func NewRenderer(tileSize, rows, cols int) *Renderer {
	return &Renderer{
		tileSize:             tileSize,
		rows:                 rows,
		cols:                 cols,
		backgroundColor:      color.RGBA{0x3e, 0x12, 0x2d, 0xff},
		boardBackgroundColor: color.RGBA{0x3e, 0x22, 0x2d, 0xff},
		boardFrameColor:      color.RGBA{0xbb, 0xad, 0xa0, 0xff},
		boardImage:           ebiten.NewImage(cols*tileSize, rows*tileSize),
		nextPieceImage:       ebiten.NewImage(4*tileSize, 4*tileSize),
		scoreImage:           ebiten.NewImage(100, 50),
		tileImage:            ebiten.NewImage(tileSize-1, tileSize-1),
	}
}

func (r *Renderer) Draw(screen *ebiten.Image, board *Board) {
	screen.Fill(r.backgroundColor)

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
	r.boardImage.Fill(r.boardBackgroundColor)
	vector.StrokeLine(r.boardImage, 0, 0, 0, float32(r.tileSize*r.rows), 1, r.boardFrameColor, true)
	vector.StrokeLine(r.boardImage, 0, 0, float32(r.tileSize*r.cols), 0, 1, r.boardFrameColor, true)
	vector.StrokeLine(r.boardImage, float32(r.tileSize*r.cols), 0, float32(r.tileSize*r.cols), float32(r.tileSize*r.rows), 1, r.boardFrameColor, true)
	vector.StrokeLine(r.boardImage, 0, float32(r.tileSize*r.rows), float32(r.tileSize*r.cols), float32(r.tileSize*r.rows), 1, r.boardFrameColor, true)

	// Tiles
	for y := 0; y < r.rows; y++ {
		for x := 0; x < r.cols; x++ {
			if board.field[y][x] == nil {
				continue
			}
			op := &ebiten.DrawImageOptions{}
			r.tileImage.Fill(board.field[y][x])
			op.GeoM.Translate(float64(x*r.tileSize)+1, float64(y*r.tileSize)+1)
			r.boardImage.DrawImage(r.tileImage, op)
		}
	}

	// Current piece
	if board.currentPiece != nil {
		for _, tile := range board.currentPiece.getTiles() {
			op := &ebiten.DrawImageOptions{}
			r.tileImage.Fill(tile.color)
			op.GeoM.Translate(
				board.currentPiece.x*float64(r.tileSize)+float64(tile.x*r.tileSize)+1,
				board.currentPiece.y*float64(r.tileSize)+float64(tile.y*r.tileSize)+1,
			)
			r.boardImage.DrawImage(r.tileImage, op)
		}
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(10, 10)
	screen.DrawImage(r.boardImage, op)
}

func (r *Renderer) renderStartGame(screen *ebiten.Image) {
	textString := "Press [Enter] to Start"
	op := &text.DrawOptions{}
	op.GeoM.Translate(float64(screenW)/2-100, float64(screenH)/2)
	op.ColorScale.ScaleWithColor(color.White)
	text.Draw(screen, textString, &text.GoTextFace{
		Source: mplusFaceSource,
		Size:   20,
	}, op)
}

func (r *Renderer) renderPauseOverlay(screen *ebiten.Image) {
	textString := "Paused"
	op := &text.DrawOptions{}
	op.GeoM.Translate(float64(screenW)/2-40, float64(screenH)/2)
	op.ColorScale.ScaleWithColor(color.White)
	text.Draw(screen, textString, &text.GoTextFace{
		Source: mplusFaceSource,
		Size:   24,
	}, op)
}

func (r *Renderer) renderGameOverOverlay(screen *ebiten.Image) {
	textString := "Game over"
	op := &text.DrawOptions{}
	op.GeoM.Translate(float64(screenW)/2-60, float64(screenH)/2-30)
	op.ColorScale.ScaleWithColor(color.White)
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

func (r *Renderer) renderNextPiece(b *Board, screen *ebiten.Image) {
	r.nextPieceImage.Fill(color.RGBA{0x20, 0x20, 0x20, 0xff})

	if len(b.pieceQueue) > 0 && b.pieceQueue[0] != nil {
		for _, tile := range b.pieceQueue[0].getTiles() {
			op := &ebiten.DrawImageOptions{}
			r.tileImage.Fill(tile.color)
			op.GeoM.Translate(
				float64((1+tile.x)*r.tileSize)+1,
				float64((1+tile.y)*r.tileSize)+1,
			)
			r.nextPieceImage.DrawImage(r.tileImage, op)
		}
	}

	opNext := &ebiten.DrawImageOptions{}
	opNext.GeoM.Translate(float64(screenW-(4*r.tileSize)-10), 10)
	screen.DrawImage(r.nextPieceImage, opNext)
}

func (r *Renderer) renderScore(b *Board, screen *ebiten.Image) {
	r.scoreImage.Clear()
	op := &text.DrawOptions{}

	text.Draw(r.scoreImage, fmt.Sprintf("Score: %d", b.Score), &text.GoTextFace{
		Source: mplusFaceSource,
		Size:   12,
	}, op)

	opScore := &ebiten.DrawImageOptions{}
	opScore.GeoM.Translate(float64(screenW-100), float64(screenH-60))

	screen.DrawImage(r.scoreImage, opScore)
}
