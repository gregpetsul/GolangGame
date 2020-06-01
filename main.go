package main

import (
	"fmt"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidth  = 800
	screenHeight = 600

	targetTicksPerSecond = 60
)

var delta float64
var scale int = 4 // for TEMPORARY testing purposes TODO: remove scale

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

	//initTileMap(renderer)
	parseMapFile(renderer, "stress_test.map")

	elements = append(elements, newPlayer(renderer))

	elements = append(elements, newEnemy(renderer, vector{300, 400}))

	initBulletPool(renderer)
	fpsTime := time.Now()
	fps := 0
	for {
		frameStartTime := time.Now()

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}

		//renderer.SetDrawColor(255, 255, 255, 255)
		renderer.Clear()
		for _, tile := range tiles {
			if tile.active {
				if getDist(tile.position, vector{screenWidth / 2, screenHeight / 2}) < 2000 {
					err = tile.update()
					if err != nil {
						fmt.Println("updating element:", err)
						return
					}
					err = tile.draw(renderer)
					if err != nil {
						fmt.Println("drawing element:", tile)
						return
					}
				}
			}
		}
		for _, elem := range elements {
			if elem.active {
				if getDist(elem.position, vector{screenWidth / 2, screenHeight / 2}) < 2000 {
					err = elem.update()
					if err != nil {
						fmt.Println("updating element:", err)
						return
					}
					err = elem.draw(renderer)
					if err != nil {
						fmt.Println("drawing element:", elem)
						return
					}
				}
			}
		}
		if err := checkCollisions(); err != nil {
			fmt.Println("checking collisions:", err)
			return
		}

		renderer.Present()

		delta = time.Since(frameStartTime).Seconds() * targetTicksPerSecond
		fps++
		if time.Since(fpsTime).Seconds() >= 1 {
			fpsTime = time.Now()
			fmt.Println(fps)
			fps = 0
		}
	}
}
