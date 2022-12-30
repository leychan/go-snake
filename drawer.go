package main

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"math"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

type Drawer struct {
	Drawer [][]string
	Body   []Location
	Header Location
	Length int
	Direct keyboard.Key
	Action chan keyboard.Key
	Food   Location
}

type Location struct {
	X int
	Y int
}

func (d *Drawer) Print() {
	for i := 0; i < d.Length; i++ {
		for j := 0; j < d.Length; j++ {
			fmt.Print(d.Drawer[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}

func (d *Drawer) Init() {
	if d.Length > math.MaxInt {
		d.Length = math.MaxInt
	}
	for i := 0; i < d.Length; i++ {
		var tmpLine []string
		for j := 0; j < d.Length; j++ {
			tmpLine = append(tmpLine, " * ")
		}
		d.Drawer = append(d.Drawer, tmpLine)
	}
}

func (d *Drawer) InitHeader() {
	d.newHeader(d.RandomLocation())
}

func (d *Drawer) IsBody(l Location) bool {
	for i := 0; i < len(d.Body); i++ {
		if d.Body[i].X == l.X && d.Body[i].Y == l.Y {
			return true
		}
	}
	return false
}

func (d *Drawer) IsHeader(l Location) bool {
	if d.Header.X == l.X && d.Header.Y == l.Y {
		return true
	}
	return false
}

func (d *Drawer) newHeader(l Location) {
	d.Header = l
	d.Body = append([]Location{l}, d.Body...)
	d.Drawer[l.X][l.Y] = " $ "
}

func (d *Drawer) removeTail() {
	d.Body = d.Body[:len(d.Body)-1]
}

func (d *Drawer) MoveDirect() {
	x := -1
	y := -1
	if d.Direct == keyboard.KeyArrowUp {
		x, y = d.Header.X+1, d.Header.Y
	} else if d.Direct == keyboard.KeyArrowDown {
		x, y = d.Header.X-1, d.Header.Y
	}
	if d.Direct == keyboard.KeyArrowRight {
		x, y = d.Header.X, d.Header.Y+1
	}
	if d.Direct == keyboard.KeyArrowLeft {
		x, y = d.Header.X, d.Header.Y-1
	}

	l := Location{x, y}
	d.die(l)
	d.newHeader(l)

	if !d.IsFood(l) {
		d.removeTail()
	} else {
		d.newFood()
	}
}

func (d *Drawer) die(l Location) {
	if l.X < 0 || l.X >= d.Length || l.Y < 0 || l.Y >= d.Length {
		fmt.Println("die")
		os.Exit(1)
	}
}

func (d *Drawer) IsFood(l Location) bool {
	if d.Food.X == l.X && d.Food.Y == l.Y {
		return true
	}

	return false
}

func (d *Drawer) SetDirect(direct keyboard.Key) {
	d.Direct = direct
}

func (d *Drawer) RandomDirect() {
	rand.Seed(time.Now().Unix())
	r := rand.Intn(3)
	if r == 0 {
		d.Direct = keyboard.KeyArrowUp
		return
	}
	if r == 1 {
		d.Direct = keyboard.KeyArrowDown
		return
	}
	if r == 2 {
		d.Direct = keyboard.KeyArrowLeft
		return
	}
	if r == 3 {
		d.Direct = keyboard.KeyArrowRight
		return
	}
}

func (d *Drawer) RandomLocation() Location {
	rand.Seed(time.Now().Unix())
	x := rand.Intn(d.Length)
	y := rand.Intn(d.Length)

	for i := 0; i < len(d.Body); i++ {
		if d.Body[i].X == x || d.Body[i].Y == y {
			d.RandomLocation()
		}
	}
	return Location{x, y}
}

func (d *Drawer) newFood() {
	d.Food = d.RandomLocation()
	d.Drawer[d.Food.X][d.Food.Y] = " @ "
}

func (d *Drawer) Clean() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
