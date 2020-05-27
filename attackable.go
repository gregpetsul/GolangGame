package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type attackable struct {
	container *element
	animator  *animator
}

func newAttackable(container *element) *attackable {
	return &attackable{
		container: container,
		animator:  container.getComponent(&animator{}).(*animator)}
}

func (atbl *attackable) onDraw(renderer *sdl.Renderer) error {
	return nil
}

func (atbl *attackable) onUpdate() error {
	if atbl.animator.finished && atbl.animator.current == "destroy" {
		atbl.container.active = false
	}
	return nil
}

func (atbl *attackable) onCollision(other *element) error {
	if other.tag == "bullet" {
		atbl.animator.setSequence("destroy")
	}
	return nil
}
