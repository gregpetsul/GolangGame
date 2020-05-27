package main

import (
	"math"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

type playerMover struct {
	container *element
	speed     float64

	//sr *spriteRenderer
}

func newPlayerMover(container *element, speed float64) *playerMover {
	return &playerMover{
		container: container,
		speed:     speed,
		//sr:        container.getComponent(&spriteRenderer{}).(*spriteRenderer),
	}
}

func (mover *playerMover) onUpdate() error {
	keys := sdl.GetKeyboardState()

	cont := mover.container

	//left and right movement
	if keys[sdl.SCANCODE_LEFT] == 1 || keys[sdl.SCANCODE_A] == 1 {
		if cont.position.x /*-(mover.sr.width/2.0)*/ > 0 {
			cont.position.x -= mover.speed * delta
		}
	} else if keys[sdl.SCANCODE_RIGHT] == 1 || keys[sdl.SCANCODE_D] == 1 {
		if cont.position.x /*+(mover.sr.height/2.0)*/ < screenWidth {
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

func (mover *playerMover) onDraw(renderer *sdl.Renderer) error {
	return nil
}
func (mover *playerMover) onCollision(other *element) error {
	return nil
}

type playerShooter struct {
	container *element
	cooldown  time.Duration
	lastShot  time.Time
}

func newPlayerShooter(container *element, cooldown time.Duration) *playerShooter {
	return &playerShooter{
		container: container,
		cooldown:  cooldown,
	}
}

func (mover *playerShooter) onUpdate() error {
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

func (mover *playerShooter) onDraw(renderer *sdl.Renderer) error {
	return nil
}
func (mover *playerShooter) onCollision(other *element) error {
	return nil
}

func (mover *playerShooter) shoot(x, y float64, xmouse, ymouse int32) {
	if bul, ok := bulletFromPool(); ok {
		bul.active = true
		bul.position.x = x
		bul.position.y = y
		bul.rotation = math.Atan2(float64(ymouse)-y, float64(xmouse)-x)
		bul.imageOffset = 45
		bul.tag = "bullet"
	}
}
