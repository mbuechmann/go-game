// +build example

package main

import (
	"git.mbuechmann.com/go-game/desktop"
	"git.mbuechmann.com/go-game/game"
	"git.mbuechmann.com/go-game/gfx"
	"git.mbuechmann.com/go-game/keys"
)

var image *gfx.Image
var renderOptions *gfx.RenderOptions

func main() {
	mode := &desktop.Mode{Width: 1280, Height: 800, Fullscreen: false}
	desktop.OpenWindow(mode)
	gfx.SetPixelSize(4)
	gfx.SetClearColor(0.5, 0.5, 0.5, 1.0)

	desktop.Run(&game.State{
		OnKeyDown:   onKeyDown,
		RenderFunc:  render,
		InitFunc:    initFunc,
		CleanupFunc: cleanup,
	})
}

func initFunc() {
	image = gfx.NewImage("assets/grey.png")
	renderOptions = gfx.NewRenderOptions()
}

func cleanup() {
	image.Delete()
}

func render() {
	gfx.Clear()

	topLimit := 300.0
	yOff := 25.0

	renderOptions.Y = 00.0
	for x := 15.0; x < topLimit; x += 30.0 {
		renderOptions.X = x
		renderOptions.A = x / topLimit
		image.Render(renderOptions)
	}

	renderOptions.Y += yOff
	for x := 15.0; x < topLimit; x += 30.0 {
		renderOptions.X = x
		renderOptions.R = x / topLimit
		image.Render(renderOptions)
	}

	renderOptions.Y += yOff
	for x := 15.0; x < topLimit; x += 30.0 {
		renderOptions.X = x
		renderOptions.G = x / topLimit
		image.Render(renderOptions)
	}

	renderOptions.Y += yOff
	for x := 15.0; x < topLimit; x += 30.0 {
		renderOptions.X = x
		renderOptions.B = x / topLimit
		image.Render(renderOptions)
	}

	renderOptions.Y += yOff
	for x := 15.0; x < topLimit; x += 30.0 {
		renderOptions.X = x
		renderOptions.R = x / topLimit
		renderOptions.G = x / topLimit
		image.Render(renderOptions)
	}

	renderOptions.Y += yOff
	for x := 15.0; x < topLimit; x += 30.0 {
		renderOptions.X = x
		renderOptions.R = x / topLimit
		renderOptions.B = x / topLimit
		image.Render(renderOptions)
	}

	renderOptions.Y += yOff
	for x := 15.0; x < topLimit; x += 30.0 {
		renderOptions.X = x
		renderOptions.G = x / topLimit
		renderOptions.B = x / topLimit
		image.Render(renderOptions)
	}

	renderOptions.Y += yOff
	for x := 15.0; x < topLimit; x += 30.0 {
		renderOptions.X = x
		renderOptions.Rot = gfx.Rotation{Angle: x * 1.2, X: 8, Y: 8}
		image.Render(renderOptions)
	}
	renderOptions.Rot = gfx.Rotation{Angle: 0, X: 8, Y: 8}
}

func onKeyDown(k keys.Key) {
	if k == keys.Esc {
		desktop.Exit()
	}
}