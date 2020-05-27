package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

//if an attackable component is added to an element, the element will be "destroyed" upon a bullet striking it
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

//will keep an "eye" on the attackable object to check if the "destroy" animation is complete then will deactivate the element
func (atbl *attackable) onUpdate() error {
	if atbl.animator.finished && atbl.animator.current == "destroy" {
		atbl.container.active = false
	}
	return nil
}

//if a collision occurs with a bullet and an attackable object set the animation sequence to its death animation
func (atbl *attackable) onCollision(other *element) error {
	if other.tag == "bullet" {
		atbl.animator.setSequence("destroy")
	}
	return nil
}
