package main

import "github.com/veandco/go-sdl2/sdl"

func newTile(renderer *sdl.Renderer, tileType string, position vector) *element {
	tile := &element{}
	tile.position = position
	tile.tag = "tile"

	tr := newTileRenderer(tile, renderer, "assets/tiles/testing/"+tileType+".bmp")
	tile.addComponent(tr)

	tile.active = true

	return tile
}

type tileRenderer struct {
	container *element
	tex       *sdl.Texture
}

func newTileRenderer(container *element, renderer *sdl.Renderer, filename string) *tileRenderer {
	tex, err := textureFromBMP(renderer, filename)
	if err != nil {
		panic(err)
	}

	return &tileRenderer{
		container: container,
		tex:       tex,
	}
}

func (tr *tileRenderer) onDraw(renderer *sdl.Renderer) error {
	return drawTexture(tr.tex, tr.container.position, 0, tr.container.size, renderer)
}
func (tr *tileRenderer) onUpdate() error {
	return nil
}
func (tr *tileRenderer) onCollision(other *element) error {
	return nil
}

func renderTile(renderer *sdl.Renderer, tileType string, position vector) {
	tile := newTile(renderer, tileType, position)
	tiles = append(tiles, tile)
}

var tiles []*element
