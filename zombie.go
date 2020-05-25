package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const basicEnemySize = 16

func newEnemy(renderer *sdl.Renderer, position vector) *element {
	enemy := &element{}
	enemy.position = position
	enemy.size = 1
	enemy.rotation = 0

	sr := newSpriteRenderer(enemy, renderer, "assets/sprites/basic_zombie.bmp")
	enemy.addComponent(sr)
	enemy.active = true

	return enemy
}
