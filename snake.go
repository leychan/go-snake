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

// InitTail Init 初始化尾部和方向
func (s *Snake) InitTail() {
	s.Tail = s.Header
}

func (s *Snake) InitDirect(h, w int) {
	s.randomDirect(h, w)
}

// IsBody 判断坐标点是不是蛇的身体
func (s *Snake) IsBody(l Location) bool {
	for i := 0; i < len(s.Body); i++ {
		if s.Body[i].X == l.X && s.Body[i].Y == l.Y {
			return true
		}
	}
	return false
}

// IsHeader 判断坐标点是不是蛇的头部
func (s *Snake) IsHeader(l Location) bool {
	if s.Header.X == l.X && s.Header.Y == l.Y {
		return true
	}
	return false
}

// NewHeader 坐标点设置为蛇新的头部
func (s *Snake) NewHeader(l Location) {
	s.Header = l
	s.Body = append([]Location{l}, s.Body...)
}

// SetDirect 设置蛇运行的方向
func (s *Snake) SetDirect(direct keyboard.Key) {
	s.Direct = direct
}

// 蛇的随机运行方向
func (s *Snake) randomDirect(h, w int) {

	if s.Header.X <= 2 && h>>2 > s.Header.X {
		s.Direct = keyboard.KeyArrowDown
		return
	}

	if s.Header.Y <= 2 && w>>2 > s.Header.Y {
		s.Direct = keyboard.KeyArrowRight
		return
	}

	if s.Header.X >= h-2 && h>>2 < s.Header.X {
		s.Direct = keyboard.KeyArrowUp
		return
	}

	if s.Header.Y >= w-2 && w>>2 < s.Header.Y {
		s.Direct = keyboard.KeyArrowLeft
		return
	}

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

// RemoveTail 更新蛇的身体和尾部坐标
func (s *Snake) RemoveTail() {
	s.Body = s.Body[:len(s.Body)-1]
	s.Tail = s.Body[len(s.Body)-1]
}

// CheckDied 检查蛇是否死掉
func (s *Snake) CheckDied(l Location, maxH int, maxW int) {
	if l.X < 0 || l.X >= maxH || l.Y < 0 || l.Y >= maxW {
		fmt.Println("hit the wall! died!!!")
		s.CurrentCond()
		os.Exit(1)
	}

	if s.IsBody(l) {
		fmt.Println("eat your self! died!!!")
		s.CurrentCond()
		os.Exit(1)
	}
}

// CurrentCond 蛇当前的情况
func (s *Snake) CurrentCond() {
	fmt.Printf("current length: %d\n", len(s.Body))
}

// NextLocation 蛇的运行方向上的下一个坐标
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
