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

	atbl := newAttackable(enemy)
	enemy.addComponent(atbl)

	col := circle{
		centre: enemy.position,
		radius: 8,
	}

	enemy.collisions = append(enemy.collisions, col)

	enemy.active = true

	return enemy
}
