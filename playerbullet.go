package main

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	bulletSize        = 8
	bulletSpeed       = 0.5
	bulletRenderCount = 100
)

type playerBullet struct {
	tex   *sdl.Texture
	x, y  float64
	angle float64

	active bool
}

func newBullet(renderer *sdl.Renderer) (bul playerBullet) {
	bul.tex = textureFromBMP(renderer, "assets/sprites/bullet_pink.bmp")
	return bul
}

func (bul *playerBullet) draw(renderer *sdl.Renderer) {
	if !bul.active {
		return
	}

	// Converting playerBullet coordinates to top left of sprite
	x := bul.x - bulletSize/2.0
	y := bul.y - bulletSize/2.0

	renderer.CopyEx(bul.tex,
		&sdl.Rect{X: 0, Y: 0, W: bulletSize, H: bulletSize},
		&sdl.Rect{X: int32(x), Y: int32(y), W: bulletSize * 2, H: bulletSize * 2},
		(bul.angle*180/math.Pi)+45,
		&sdl.Point{X: bulletSize, Y: bulletSize},
		sdl.FLIP_NONE)
}

func (bul *playerBullet) update() {
	bul.x += bulletSpeed * math.Cos(bul.angle)
	bul.y += bulletSpeed * math.Sin(bul.angle)

	if bul.x > screenWidth || bul.x < 0 || bul.y > screenHeight || bul.y < 0 {
		bul.active = false
	}
}

var bulletPool []*playerBullet

func initBulletPool(renderer *sdl.Renderer) {
	for i := 0; i < bulletRenderCount; i++ {
		bul := newBullet(renderer)
		bulletPool = append(bulletPool, &bul)
	}
}

//try to grab a playerBullet from the playerBullet pool(only non-active bullets)
func bulletFromPool() (*playerBullet, bool) {
	for _, bul := range bulletPool {
		if !bul.active {
			return bul, true
		}
	}

	return nil, false
}
