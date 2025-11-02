package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	pressDelayTicks          = 10 // 1/6 of second
	pressRepeatIntervalTicks = 2 // 1/30 of second
)

type InputHandler struct{}

func (i *InputHandler) Update(board *Board) {
	if board == nil {
		return
	}

	// Handle pause/unpause toggle first.
	if inpututil.IsKeyJustPressed(ebiten.KeyP) {
		board.TogglePause()
	}

	// If the game is stopped (paused or game over), don't process any other input.
	if board.isStopped() {
		return
	}

	if keyPressAndMove(ebiten.KeyArrowLeft) || keyPressAndMove(ebiten.KeyA) {
		board.MoveLeft()
	}

	if keyPressAndMove(ebiten.KeyArrowRight) || keyPressAndMove(ebiten.KeyD) {
		board.MoveRight()
	}

	if keyPressAndMove(ebiten.KeyArrowDown) || keyPressAndMove(ebiten.KeyS) {
		board.MoveDown()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		board.Fall()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) || inpututil.IsKeyJustPressed(ebiten.KeyW) {
		board.Rotate()
	}
}

func keyPressAndMove(key ebiten.Key) bool {
	return inpututil.IsKeyJustPressed(key) ||
		(inpututil.KeyPressDuration(key) > pressDelayTicks &&
			inpututil.KeyPressDuration(key)%pressRepeatIntervalTicks == 0)
}
