package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const basicEnemySize = 16

func newEnemy(renderer *sdl.Renderer, position vector) *element {
	enemy := &element{}
	enemy.position = position
	enemy.size = 1
	enemy.rotation = 0

	// sr := newSpriteRenderer(enemy, renderer, "assets/sprites/basic_zombie.bmp")
	// enemy.addComponent(sr)
	idleSequence, err := newSequence("assets/sprites/zombie/idle", 5, true, renderer)
	if err != nil {
		panic(fmt.Errorf("creating idle sequence: %v", err))
	}
	destroySequence, err := newSequence("assets/sprites/zombie/destroy", 5, false, renderer)
	if err != nil {
		panic(fmt.Errorf("creating destroy sequence: %v", err))
	}
	sequences := map[string]*sequence{
		"idle":    idleSequence,
		"destroy": destroySequence,
	}
	animator := newAnimator(enemy, sequences, "idle")
	enemy.addComponent(animator)

	atbl := newAttackable(enemy)
	enemy.addComponent(atbl)

	col := circle{
		centre: enemy.position,
		radius: 8 * float64(scale), //TODO: remove scale
	}

	enemy.collisions = append(enemy.collisions, col)

	enemy.active = true

	return enemy
}
