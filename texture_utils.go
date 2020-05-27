package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

func drawTexture(tex *sdl.Texture, position vector, rotation float64, size float64, renderer *sdl.Renderer) error {

	_, _, width, height, err := tex.Query()
	if err != nil {
		return fmt.Errorf("querying texture: %v", err)
	}
	// Converting coordinates to top left of sprite
	x := position.x - float64(width)/2.0
	y := position.y - float64(height)/2.0

	if size == 0 {
		size = 1
	}

	return renderer.CopyEx(
		tex,
		&sdl.Rect{X: 0, Y: 0, W: int32(width), H: int32(height)},
		&sdl.Rect{X: int32(x), Y: int32(y), W: int32(width) * int32(size) * scale, H: int32(height) * int32(size) * scale}, //TODO: remove scale (twice)
		rotation,
		&sdl.Point{X: int32(width) / 2, Y: int32(height) / 2},
		sdl.FLIP_NONE)
}

func textureFromBMP(renderer *sdl.Renderer, filename string) (*sdl.Texture, error) {
	img, err := sdl.LoadBMP(filename)
	if err != nil {
		return nil, fmt.Errorf("loading %v: %v", filename, err)
	}
	defer img.Free()
	tex, err := renderer.CreateTextureFromSurface(img)
	if err != nil {
		return nil, fmt.Errorf("creating texture from %v: %v", filename, err)
	}

	return tex, nil
}
