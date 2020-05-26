package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type attackable struct {
	container *element
}

func newAttackable(container *element) *attackable {
	return &attackable{container: container}
}

func (atbl *attackable) onDraw(renderer *sdl.Renderer) error {
	return nil
}

func (atbl *attackable) onUpdate() error {
	return nil
}

func (atbl *attackable) onCollision(other *element) error {
	if other.tag == "bullet" {
		atbl.container.active = false
	}
	return nil
}
