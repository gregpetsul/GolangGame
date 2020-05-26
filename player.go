package main

import (
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	playerSpeed        = 5 //pixels per second
	playerSize         = 16
	playerShotCooldown = time.Millisecond * 500
)

func newPlayer(renderer *sdl.Renderer) *element {
	player := &element{}

	player.position = vector{
		x: 250,
		y: 250,
	}

	player.active = true

	sr := newSpriteRenderer(player, renderer, "assets/sprites/player_pink.bmp")
	player.addComponent(sr)

	mover := newKeyboardMover(player, playerSpeed)
	player.addComponent(mover)

	shooter := newKeyboardShooter(player, playerShotCooldown)
	player.addComponent(shooter)

	return player
}
