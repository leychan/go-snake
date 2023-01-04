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
	MaxH    int
	MaxW    int
}

// InitTail Init 初始化尾部和方向
func (s *Snake) InitTail() {
	s.Tail = s.Header
}

func (s *Snake) InitDirect() {
	s.randomDirect()
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
func (s *Snake) randomDirect() {

	if s.Header.X <= 2 && s.MaxH>>2 > s.Header.X {
		s.Direct = keyboard.KeyArrowDown
		return
	}

	if s.Header.Y <= 2 && s.MaxH>>2 > s.Header.Y {
		s.Direct = keyboard.KeyArrowRight
		return
	}

	if s.Header.X >= s.MaxH-2 && s.MaxH>>2 < s.Header.X {
		s.Direct = keyboard.KeyArrowUp
		return
	}

	if s.Header.Y >= s.MaxW-2 && s.MaxW>>2 < s.Header.Y {
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
func (s *Snake) CheckDied(l Location) {
	if l.X < 0 || l.X >= s.MaxH || l.Y < 0 || l.Y >= s.MaxW {
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

// Run 蛇运行,每次蛇头的位置都会变换到新位置
func (s *Snake) Run() {
	//贪吃蛇蛇头的下一个位置
	l := s.NextLocation()

	//判断下一个位置是否是正常的位置,如果是撞墙或者是吃到自己的身体,则死亡,游戏结束
	s.CheckDied(l)

	//新位置成为蛇头
	s.NewHeader(l)
}
