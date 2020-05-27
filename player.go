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

	player.velocity = vector{
		x: 0,
		y: 0,
	}

	player.active = true

	frontIdleSequence, err := newSequence("assets/sprites/player/front_idle", 5, true, renderer)
	if err != nil {
		panic(fmt.Errorf("creating front_idle sequence: %v", err))
	}
	frontWalkSequence, err := newSequence("assets/sprites/player/front_walk", playerSpeed, true, renderer)
	if err != nil {
		panic(fmt.Errorf("creating front_walk sequence: %v", err))
	}
	backIdleSequence, err := newSequence("assets/sprites/player/back_idle", 5, true, renderer)
	if err != nil {
		panic(fmt.Errorf("creating back_idle sequence: %v", err))
	}
	backWalkSequence, err := newSequence("assets/sprites/player/back_walk", playerSpeed, true, renderer)
	if err != nil {
		panic(fmt.Errorf("creating back_walk sequence: %v", err))
	}

	sequences := map[string]*sequence{
		"front_idle": frontIdleSequence,
		"front_walk": frontWalkSequence,
		"back_idle":  backIdleSequence,
		"back_walk":  backWalkSequence,
	}
	animator := newAnimator(player, sequences, "front_idle")
	player.addComponent(animator)

	mover := newPlayerMover(player, playerSpeed)
	player.addComponent(mover)

	shooter := newPlayerShooter(player, playerShotCooldown)
	player.addComponent(shooter)

	return player
}
