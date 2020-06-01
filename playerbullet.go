package main

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	bulletSize        = 8
	bulletSpeed       = 20 //pixels per second
	bulletRenderCount = 25
)

func newBullet(renderer *sdl.Renderer) *element {
	bullet := &element{}

	sr := newSpriteRenderer(bullet, renderer, "assets/sprites/bullet_pink.bmp")
	bullet.addComponent(sr)

	mover := newBulletMover(bullet, bulletSpeed)
	bullet.addComponent(mover)

	col := circle{
		centre: bullet.position,
		radius: 2 * float64(scale), //TODO: remove scale
	}
	bullet.collisions = append(bullet.collisions, col)

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

func (mover *bulletMover) onCollision(other *element) error {
	mover.container.active = false
	return nil
}

func (mover *bulletMover) onUpdate() error {
	c := mover.container

	c.position.x += bulletSpeed * math.Cos(c.rotation) * delta
	c.position.y += bulletSpeed * math.Sin(c.rotation) * delta

	if c.position.x > screenWidth || c.position.x < 0 ||
		c.position.y > screenHeight || c.position.y < 0 {
		c.active = false
	}

	c.collisions[0].centre = c.position

	return nil
}
