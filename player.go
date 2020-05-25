package main

import (
	"math"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	playerSpeed        = 0.05
	playerSize         = 16
	playerShotCooldown = time.Millisecond * 100
)

type player struct {
	tex      *sdl.Texture
	x, y     float64
	lastShot time.Time
}

func newPlayer(renderer *sdl.Renderer) (p player) {
	p.tex = textureFromBMP(renderer, "assets/sprites/player_pink.bmp")

	p.x = screenWidth / 2.0
	p.y = screenHeight / 2.0

	return p
}

func (p *player) draw(renderer *sdl.Renderer) {
	//converting player coordinates to top left
	x := p.x - playerSize/2.0
	y := p.y - playerSize/2.0
	renderer.Copy(p.tex,
		&sdl.Rect{X: 0, Y: 0, W: playerSize, H: playerSize},
		&sdl.Rect{X: int32(x), Y: int32(y), W: playerSize * 2, H: playerSize * 2})
}

func (p *player) update() {
	keys := sdl.GetKeyboardState()
	xmouse, ymouse, mouse := sdl.GetMouseState()
	//left and right movement
	if keys[sdl.SCANCODE_LEFT] == 1 || keys[sdl.SCANCODE_A] == 1 {
		p.x -= playerSpeed
	} else if keys[sdl.SCANCODE_RIGHT] == 1 || keys[sdl.SCANCODE_D] == 1 {
		p.x += playerSpeed
	}
	//up and down movement
	if keys[sdl.SCANCODE_DOWN] == 1 || keys[sdl.SCANCODE_S] == 1 {
		p.y += playerSpeed
	} else if keys[sdl.SCANCODE_UP] == 1 || keys[sdl.SCANCODE_W] == 1 {
		p.y -= playerSpeed
	}

	if mouse == 1 {
		if time.Since(p.lastShot) >= playerShotCooldown {
			p.shoot(p.x, p.y, xmouse, ymouse)

			p.lastShot = time.Now()
		}
	}
}

func (p *player) shoot(xplayer, yplayer float64, xmouse, ymouse int32) {
	if bul, ok := bulletFromPool(); ok {
		bul.active = true
		bul.x = xplayer
		bul.y = yplayer
		bul.angle = math.Atan2(float64(ymouse)-yplayer, float64(xmouse)-xplayer)
	}
}
