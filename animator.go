package main

import (
	"fmt"
	"io/ioutil"
	"path"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

type animator struct {
	container      *element
	sequences      map[string]*sequence
	current        string
	lastFramChange time.Time
	finished       bool
}

func newAnimator(container *element, sequences map[string]*sequence, defaultSequence string) *animator {
	var an animator
	an.container = container
	an.sequences = sequences
	an.current = defaultSequence
	an.lastFramChange = time.Now()

	return &an
}

func (an *animator) onUpdate() error {
	sequence := an.sequences[an.current]
	frameInterval := float64(time.Second) / sequence.sampleRate
	if time.Since(an.lastFramChange) >= time.Duration(frameInterval) {
		an.finished = sequence.nextFrame()
		an.lastFramChange = time.Now()
	}
	return nil
}

func (an *animator) onDraw(renderer *sdl.Renderer) error {
	tex := an.sequences[an.current].texture()

	return drawTexture(tex, an.container.position, 0, an.container.rotation, renderer)
}

func (an *animator) onCollision(other *element) error {
	return nil
}

func (an *animator) setSequence(name string) {
	if an.current != name {
		an.current = name
		an.lastFramChange = time.Now()
	}
}

type sequence struct {
	textures   []*sdl.Texture
	frame      int
	sampleRate float64
	loop       bool
}

func newSequence(filepath string, sampleRate float64, loop bool, renderer *sdl.Renderer) (*sequence, error) {
	var seq sequence
	files, err := ioutil.ReadDir(filepath)
	if err != nil {
		return nil, fmt.Errorf("reading directory %v: %v", filepath, err)
	}

	for _, file := range files {
		filename := path.Join(filepath, file.Name())
		tex, err := textureFromBMP(renderer, filename)
		if err != nil {
			return nil, fmt.Errorf("loading sequence frame: %v", err)
		}
		seq.textures = append(seq.textures, tex)
	}
	seq.sampleRate = sampleRate
	seq.loop = loop

	return &seq, nil
}

func (seq *sequence) texture() *sdl.Texture {
	return seq.textures[seq.frame]
}

//gets the next frame and indicates whether or no animation is complete
func (seq *sequence) nextFrame() bool {
	if seq.frame == len(seq.textures)-1 {
		if seq.loop { //reset loop if it is a looping animation
			seq.frame = 0
		} else {
			return true
		}
	} else {
		seq.frame++
	}
	return false
}
