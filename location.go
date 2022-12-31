package main

import (
	"math/rand"
	"time"
)

type Location struct {
	X int
	Y int
}

func RandomLocation(h, w int) Location {
	rand.Seed(time.Now().Unix())
	x := rand.Intn(h)
	y := rand.Intn(w)

	return Location{x, y}
}

func RandomLocationWithExtra(h int, w int, locations []Location) Location {
	rand.Seed(time.Now().Unix())
	x := rand.Intn(h)
	y := rand.Intn(w)
	for i := 0; i < len(locations); i++ {
		if x == locations[i].X && y == locations[i].Y {
			return RandomLocationWithExtra(h, w, locations)
		}
	}
	return Location{x, y}
}
