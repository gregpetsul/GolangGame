package main

import (
	"math"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

type keyboardMover struct {
	container *element
	speed     float64

	sr *spriteRenderer
}

func newKeyboardMover(container *element, speed float64) *keyboardMover {
	return &keyboardMover{
		container: container,
		speed:     speed,
		sr:        container.getComponent(&spriteRenderer{}).(*spriteRenderer),
	}
}

func (mover *keyboardMover) onUpdate() error {
	keys := sdl.GetKeyboardState()

	cont := mover.container

	//left and right movement
	if keys[sdl.SCANCODE_LEFT] == 1 || keys[sdl.SCANCODE_A] == 1 {
		if cont.position.x-(mover.sr.width/2.0) > 0 {
			cont.position.x -= mover.speed * delta
		}
	} else if keys[sdl.SCANCODE_RIGHT] == 1 || keys[sdl.SCANCODE_D] == 1 {
		if cont.position.x+(mover.sr.height/2.0) < screenWidth {
			cont.position.x += mover.speed * delta
		}
	}
	//up and down movement
	if keys[sdl.SCANCODE_DOWN] == 1 || keys[sdl.SCANCODE_S] == 1 {
		cont.position.y += mover.speed * delta
	} else if keys[sdl.SCANCODE_UP] == 1 || keys[sdl.SCANCODE_W] == 1 {
		cont.position.y -= mover.speed * delta
	}

	return nil
}

func (mover *keyboardMover) onDraw(renderer *sdl.Renderer) error {
	return nil
}
func (mover *keyboardMover) onCollision(other *element) error {
	return nil
}

type keyboardShooter struct {
	container *element
	cooldown  time.Duration
	lastShot  time.Time
}

func newKeyboardShooter(container *element, cooldown time.Duration) *keyboardShooter {
	return &keyboardShooter{
		container: container,
		cooldown:  cooldown,
	}
}

func (mover *keyboardShooter) onUpdate() error {
	xmouse, ymouse, mouse := sdl.GetMouseState()

	pos := mover.container.position

	if mouse == 1 {
		if time.Since(mover.lastShot) >= mover.cooldown {
			mover.shoot(pos.x, pos.y, xmouse, ymouse)

			mover.lastShot = time.Now()
		}
	}

	return nil
}

func (mover *keyboardShooter) onDraw(renderer *sdl.Renderer) error {
	return nil
}
func (mover *keyboardShooter) onCollision(other *element) error {
	return nil
}

func (mover *keyboardShooter) shoot(x, y float64, xmouse, ymouse int32) {
	if bul, ok := bulletFromPool(); ok {
		bul.active = true
		bul.position.x = x
		bul.position.y = y
		bul.rotation = math.Atan2(float64(ymouse)-y, float64(xmouse)-x)
		bul.imageOffset = 45
		bul.tag = "bullet"
	}
}
