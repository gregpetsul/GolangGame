package main

import "math"

type circle struct {
	centre vector
	radius float64
}

func collides(c1, c2 circle) bool {
	dist := math.Sqrt(math.Pow(c2.centre.x-c1.centre.x, 2) +
		math.Pow(c2.centre.y-c1.centre.y, 2))
	return dist <= c1.radius+c2.radius
}
