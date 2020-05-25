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

// type playerBullet struct {
// 	tex   *sdl.Texture
// 	x, y  float64
// 	angle float64

// 	active bool
// }

func newBullet(renderer *sdl.Renderer) *element {
	bullet := &element{}

	sr := newSpriteRenderer(bullet, renderer, "assets/sprites/bullet_pink.bmp")
	bullet.addComponent(sr)

	mover := newBulletMover(bullet, bulletSpeed)
	bullet.addComponent(mover)

	bullet.active = false

	return bullet
}

var bulletPool []*element

func initBulletPool(renderer *sdl.Renderer) {
	for i := 0; i < bulletRenderCount; i++ {
		bul := newBullet(renderer)
		elements = append(elements, bul)
		bulletPool = append(bulletPool, bul)
	}
}

func bulletFromPool() (*element, bool) {
	for _, bul := range bulletPool {
		if !bul.active {
			return bul, true
		}
	}

	return nil, false
}

type bulletMover struct {
	container *element
	speed     float64
}

func newBulletMover(container *element, speed float64) *bulletMover {
	return &bulletMover{
		container: container,
		speed:     speed,
	}
}

func (mover *bulletMover) onDraw(renderer *sdl.Renderer) error {
	return nil
}

func (mover *bulletMover) onUpdate() error {
	cont := mover.container

	cont.position.x += bulletSpeed * math.Cos(cont.rotation)
	cont.position.y += bulletSpeed * math.Sin(cont.rotation)

	if cont.position.x > screenWidth || cont.position.x < 0 ||
		cont.position.y > screenHeight || cont.position.y < 0 {
		cont.active = false
	}

	return nil
}
