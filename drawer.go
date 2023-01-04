package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

type Drawer struct {
	Drawer   [][]string
	Height   int
	Width    int
	Food     Location
	Interval time.Duration
	Snake    *Snake
	Unused   []Location
}

const FoodSTR = " @ "
const SnakeBody = " * "
const DrawerSTR = " - "

func (d *Drawer) Print() {
	for i := 0; i < d.Height; i++ {
		for j := 0; j < d.Width; j++ {
			fmt.Print(d.Drawer[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}

// Init 初始化
func (d *Drawer) Init() {
	//画布大小限制
	if d.Height > math.MaxInt {
		d.Width = math.MaxInt
	}

	//初始化画布
	for i := 0; i < d.Height; i++ {
		var tmpLine []string
		for j := 0; j < d.Width; j++ {
			tmpLine = append(tmpLine, DrawerSTR)
			d.Unused = append(d.Unused, Location{i, j})
		}
		d.Drawer = append(d.Drawer, tmpLine)
	}

	d.Clear()
	d.Print()

	//初始画布刷新时间为800ms
	d.Interval = time.Millisecond * 800

	//初始化蛇结构
	d.Snake = &Snake{}
	//初始化蛇
	l := d.RandomUnusedLocation()
	d.newSnakeHeader(l)
	d.Snake.InitTail()
	d.Snake.InitDirect(d.Height, d.Width)

	//蛇的身体放入画布已经使用的位置集合
	d.removeFromUnused(d.Snake.Header)
	//生成第一个食物
	d.newFood()
}

// RefreshLocation 刷新当前位置的元素
func (d *Drawer) RefreshLocation(l Location, c string) {
	cmdStr := fmt.Sprintf(`tput civis && tput sc && tput cup %d %d && echo "%s" && tput rc`, l.X, l.Y*3, c)
	cmd := exec.Command("/bin/bash", "-c", cmdStr)
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}

// RemoveSnakeTail 去掉当前贪吃蛇的尾部
func (d *Drawer) RemoveSnakeTail() {
	d.RefreshLocation(d.Snake.Tail, DrawerSTR)
	d.Unused = append(d.Unused, d.Snake.Tail)
	d.Snake.RemoveTail()
}

// 新位置更新为蛇头
func (d *Drawer) newSnakeHeader(l Location) {
	d.removeFromUnused(l)
	d.Snake.NewHeader(l)
	d.RefreshLocation(l, SnakeBody)
}

// Clear 清除当前屏幕
func (d *Drawer) Clear() {
	cmd := exec.Command("/bin/bash", "-c", "tput cup 0 0 && tput ed")
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}

// Run 运行
func (d *Drawer) Run() {
	//贪吃蛇蛇头的下一个位置
	l := d.Snake.NextLocation()

	//判断下一个位置是否是正常的位置,如果是撞墙或者是吃到自己的身体,则死亡,游戏结束
	d.Snake.CheckDied(l, d.Height, d.Width)

	//新位置成为蛇头
	d.newSnakeHeader(l)

	//如果下一个位置是食物,则加入身体,蛇的尾部不用去除
	//否则需要去掉蛇身体的最后一个元素
	if !d.IsFood(l) {
		d.RemoveSnakeTail()
	} else {
		//蛇吃掉食物以后,需要新生成一个食物
		d.newFood()
		//蛇吃掉食物以后,自身长度增加,运行时间间隔需要减小,以加快蛇运行的速度
		d.refreshInterval()
	}
}

// IsFood 判断当前位置是否是食物
func (d *Drawer) IsFood(l Location) bool {
	if d.Food.X == l.X && d.Food.Y == l.Y {
		return true
	}

	return false
}

// 生成食物
func (d *Drawer) newFood() {
	d.Food = d.RandomUnusedLocation()
	d.removeFromUnused(d.Food)
	d.RefreshLocation(d.Food, FoodSTR)
}

// 根据算法更新画布的刷新时间
func (d *Drawer) refreshInterval() {
	sub := time.Duration(0)
	if d.Interval > 400*time.Millisecond {
		sub = 50
	} else if d.Interval > 200*time.Millisecond {
		sub = 25
	} else {
		sub = 10
	}
	d.Interval = d.Interval - sub*time.Millisecond
	if d.Interval < 100*time.Millisecond {
		d.Interval = 100 * time.Millisecond
	}
}

// 从位置用的坐标集合中删除指定坐标
func (d *Drawer) removeFromUnused(l Location) {
	for i := 0; i < len(d.Unused); i++ {
		if d.Unused[i].X == l.X && d.Unused[i].Y == l.Y {
			d.Unused = append(d.Unused[:i], d.Unused[i:]...)
			break
		}
	}
}

// RandomUnusedLocation 从未使用过的坐标集合中随机选取一个位置
func (d *Drawer) RandomUnusedLocation() Location {
	rand.Seed(time.Now().UnixMicro())
	randLoc := rand.Intn(len(d.Unused))
	loc := d.Unused[randLoc]
	d.Unused = append(d.Unused[:randLoc], d.Unused[randLoc:]...)
	return loc
}
