package main

import (
	"fmt"
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
		x: screenWidth / 2,
		y: screenHeight / 2,
	}

	player.active = true

	idleSequence, err := newSequence("assets/sprites/player/idle", 5, true, renderer)
	if err != nil {
		panic(fmt.Errorf("creating idle sequence: %v", err))
	}
	sequences := map[string]*sequence{
		"idle": idleSequence,
	}
	animator := newAnimator(player, sequences, "idle")
	player.addComponent(animator)

	// sr := newSpriteRenderer(player, renderer, "assets/sprites/player_pink.bmp")
	// player.addComponent(sr)

	mover := newPlayerMover(player, playerSpeed)
	player.addComponent(mover)

	shooter := newPlayerShooter(player, playerShotCooldown)
	player.addComponent(shooter)

	return player
}
