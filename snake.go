package main

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"math/rand"
	"os"
	"time"
)

type Snake struct {
	Body    []Location
	Header  Location
	Direct  keyboard.Key
	Tail    Location
	BodyStr string
}

func (s *Snake) Init(h, w int) {
	l := RandomLocation(h, w)
	s.NewHeader(l)
	s.RandomDirect()
}

func (s *Snake) IsBody(l Location) bool {
	for i := 0; i < len(s.Body); i++ {
		if s.Body[i].X == l.X && s.Body[i].Y == l.Y {
			return true
		}
	}
	return false
}

func (s *Snake) IsHeader(l Location) bool {
	if s.Header.X == l.X && s.Header.Y == l.Y {
		return true
	}
	return false
}

func (s *Snake) NewHeader(l Location) {
	s.Header = l
	s.Body = append([]Location{l}, s.Body...)
}

func (s *Snake) SetDirect(direct keyboard.Key) {
	s.Direct = direct
}

func (s *Snake) RandomDirect() {
	rand.Seed(time.Now().Unix())
	r := rand.Intn(3)
	if r == 0 {
		s.Direct = keyboard.KeyArrowUp
		return
	}
	if r == 1 {
		s.Direct = keyboard.KeyArrowDown
		return
	}
	if r == 2 {
		s.Direct = keyboard.KeyArrowLeft
		return
	}
	if r == 3 {
		s.Direct = keyboard.KeyArrowRight
		return
	}
}

func (s *Snake) RemoveTail() {
	s.Body = s.Body[:len(s.Body)-1]
}

func (s *Snake) CheckDied(l Location, maxH int, maxW int) {
	if l.X < 0 || l.X >= maxH || l.Y < 0 || l.Y >= maxW {
		fmt.Println("hit the wall! died!!!")
		os.Exit(1)
	}

	if s.IsBody(l) {
		fmt.Println("eat your self! died!!!")
		os.Exit(1)
	}
}

func (s *Snake) NextLocation() Location {
	x := -1
	y := -1
	if s.Direct == keyboard.KeyArrowUp {
		x, y = s.Header.X-1, s.Header.Y
	} else if s.Direct == keyboard.KeyArrowDown {
		x, y = s.Header.X+1, s.Header.Y
	} else if s.Direct == keyboard.KeyArrowRight {
		x, y = s.Header.X, s.Header.Y+1
	} else if s.Direct == keyboard.KeyArrowLeft {
		x, y = s.Header.X, s.Header.Y-1
	}
	return Location{x, y}
}
