// +build example

package main

import (
	"github.com/go-game/go-game/desktop"
	"github.com/go-game/go-game/game"
	"github.com/go-game/go-game/gfx"
	"github.com/go-game/go-game/keys"
)

const (
	Y     = 100
	X     = 100
	WIDTH = 20
)

var height int32

func main() {
	mode := &desktop.Mode{Width: 1280, Height: 800, Fullscreen: false}
	window := desktop.OpenWindow(mode)
	gfx.SetPixelSize(4)

	window.Run(&game.State{
		OnRender:     onRender,
		OnKeyDown:    onKeyDown,
		OnMouseWheel: onMouseWheel,
	})
}

func onMouseWheel(x, y int32) {
	height += y * 2
}

func onRender() {
	gfx.Clear()
	gfx.RenderRectangle(false, X, Y, X+WIDTH, float64(Y+height))
}

func onKeyDown(k keys.Key) {
	if k == keys.Esc {
		desktop.Exit()
	}
}
