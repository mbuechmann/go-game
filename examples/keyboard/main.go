package main

import (
	"time"

	"git.mbuechmann.com/go-game/base"
	"git.mbuechmann.com/go-game/gfx"
	"github.com/go-gl/glfw/v3.1/glfw"

	"git.mbuechmann.com/go-game/examples/keyboard/sprites"
)

var heart *sprites.Heart
var vX float32
var vY float32

func main() {
	gameState := &base.GameState{
		InitFunc:    initGame,
		RenderFunc:  render,
		UpdateFunc:  logic,
		CleanupFunc: cleanupGame,
		KeyHandler:  onKey,
	}

	game := &base.Game{
		GameState: gameState,
		PixelSize: 2,
		Title:     "Keyboard-Movement",
	}
	game.Run()
}

func logic(elapsed time.Duration) {
	heart.Update(float64(elapsed))
}

func render() {
	gfx.Clear()

	heart.Render()
}

func initGame() {
	heart = sprites.NewHeart()
}

func cleanupGame() {
	heart.Delete()
}

func onKey(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if key == glfw.KeyEscape && action == glfw.Press {
		w.SetShouldClose(true)
	}

	if key == glfw.KeyUp {
		if action == glfw.Press {
			vY++
		}
		if action == glfw.Release {
			vY--
		}
	}
	if key == glfw.KeyDown {
		if action == glfw.Press {
			vY--
		}
		if action == glfw.Release {
			vY++
		}
	}
	if key == glfw.KeyRight {
		if action == glfw.Press {
			vX++
		}
		if action == glfw.Release {
			vX--
		}
	}
	if key == glfw.KeyLeft {
		if action == glfw.Press {
			vX--
		}
		if action == glfw.Release {
			vX++
		}
	}

	heart.SetDirection(vX, vY)
}
