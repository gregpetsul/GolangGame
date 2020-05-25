package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidth  = 1600
	screenHeight = 900
)

func textureFromBMP(renderer *sdl.Renderer, filename string) *sdl.Texture {
	img, err := sdl.LoadBMP(filename)
	if err != nil {
		panic(fmt.Errorf("loading %v: %v", filename, err))
	}
	defer img.Free()
	tex, err := renderer.CreateTextureFromSurface(img)
	if err != nil {
		panic(fmt.Errorf("creating texture from %v: %v", filename, err))
	}

	return tex
}

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Println("initializing SDL:", err)
		return
	}

	window, err := sdl.CreateWindow(
		"game name",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		screenWidth, screenHeight,
		sdl.WINDOW_OPENGL)
	if err != nil {
		fmt.Println("initializing window:", err)
		return
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("initializing renderer:", err)
		return
	}
	defer renderer.Destroy()

	plr := newPlayer(renderer)

	enemy := newBasicEnemy(renderer, 300, 450)

	initBulletPool(renderer)

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}

		renderer.SetDrawColor(255, 255, 255, 255)
		renderer.Clear()

		plr.draw(renderer)
		plr.update()

		enemy.draw(renderer)

		for _, bul := range bulletPool {
			bul.draw(renderer)
			bul.update()
		}

		renderer.Present()
	}
}
