package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	d = 100
	w = 119
	g = 103
)

/*
takes in a file where each character represents a different tile:
e.g.
wwwwwwwwwwwwwww
wgggdgggggggggw
wgggdgggggggggw
wgggddggggggggw
wggggdddddddddw
wggggdggggggggw
wggggdggggggggw
wggggdddddggggw
wggggggggdggggw
wggdddddddggggw
wggdggggggggggw
wggdggggggggggw
wggdggggggggggw
wggdggggggggggw
wwwwwwwwwwwwwww

then renders the appropriate tile in the relative position on the players screen
*/
func parseMapFile(renderer *sdl.Renderer, fileName string) {
	file, err := os.Open("maps/" + fileName)
	if err != nil {
		panic(fmt.Errorf("openning map file: %v", err))
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	row := 0
	var newLine string
	for scanner.Scan() {

		//first remove any
		for _, c := range scanner.Text() {
			switch c {
			case d:
				newLine += string(c)
			case w:
				newLine += string(c)
			case g:
				newLine += string(c)
			default:
				fmt.Printf("cannot render character %s. skipping over it", string(c))
			}
		}
		for col, c := range newLine {
			switch c {
			case d:
				renderTile(renderer, "dirt", vector{x: float64(col * 16 * scale), y: float64(row * 16 * scale)})
			case w:
				renderTile(renderer, "water", vector{x: float64(col * 16 * scale), y: float64(row * 16 * scale)})
			case g:
				renderTile(renderer, "grass", vector{x: float64(col * 16 * scale), y: float64(row * 16 * scale)})
			}

		}
		newLine = ""
		row++
	}
	if err := scanner.Err(); err != nil {
		panic(fmt.Errorf("reading map file lines: %v", err))
	}

}
