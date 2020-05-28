package main

import "github.com/veandco/go-sdl2/sdl"

func newTile(renderer *sdl.Renderer, tileType string, position vector) *element {
	tile := &element{}
	tile.position = position
	tile.tag = "tile"

	sr := newSpriteRenderer(tile, renderer, "assets/tiles/testing/"+tileType+".bmp")
	tile.addComponent(sr)

	tile.active = true

	return tile
}

// func initTileMap(renderer *sdl.Renderer, m [][]string) {
// 	for x := 0; x <= 15; x++ {
// 		for y := 0; y <= 15; y++ {
// 			tile := newTile(renderer, "dirt", vector{x: float64(x*16) * float64(scale), y: float64(y*16) * float64(scale)})
// 			elements = append(elements, tile)
// 		}
// 	}
//}

func renderTile(renderer *sdl.Renderer, tileType string, position vector) {
	tile := newTile(renderer, tileType, position)
	elements = append(elements, tile)
}
