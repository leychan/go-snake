package main

import (
	"fmt"
	"math"
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
	Used     []Location
}

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
			tmpLine = append(tmpLine, " * ")
		}
		d.Drawer = append(d.Drawer, tmpLine)
	}

	//初始画布刷新时间为1s
	d.Interval = time.Second

	//初始化蛇结构
	d.Snake = &Snake{}
	//初始化蛇的运行方向
	d.Snake.RandomDirect()
	//初始化蛇
	d.Snake.Init(d.Height, d.Width)
	//蛇的身体放入画布已经使用的位置集合
	d.AddUsed(d.Snake.Header, " $ ")
	//生成第一个食物
	d.newFood()
	d.Print()
}

// AddUsed 添加位置到已使用的位置集合中
func (d *Drawer) AddUsed(l Location, c string) {
	d.Used = append(d.Used, l)
	d.Drawer[l.X][l.Y] = c
}

// Run 运行
func (d *Drawer) Run() {
	//贪吃蛇蛇头的下一个位置
	l := d.Snake.NextLocation()

	//判断下一个位置是否是正常的位置,如果是撞墙或者是吃到自己的身体,则死亡,游戏结束
	d.Snake.CheckDied(l, d.Height, d.Width)

	//新位置成为蛇头
	d.Snake.NewHeader(l)
	//画布蛇身体相应位置的字符更新为` $ `
	d.Drawer[l.X][l.Y] = " $ "

	//如果下一个位置是食物,则加入身体,蛇的尾部不用去除
	//否则需要去掉蛇身体的最后一个元素
	if !d.IsFood(l) {
		d.Drawer[d.Snake.Body[len(d.Snake.Body)-1].X][d.Snake.Body[len(d.Snake.Body)-1].Y] = " * "
		d.Snake.RemoveTail()
	} else {
		//蛇吃掉食物以后,需要新生成一个食物
		d.newFood()
		//蛇吃掉食物以后,自身长度增加,运行时间间隔需要减小,以加快蛇运行的速度
		d.refreshInterval()
	}
	//每运动一步,需要刷新画布,达到运行的视觉效果
	d.Refresh()
}

// Refresh 清屏,然后打印新的画布内容
func (d *Drawer) Refresh() {
	d.Clean()
	d.Print()
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
	d.Food = RandomLocationWithExtra(d.Height, d.Width, d.Used)
	d.AddUsed(d.Food, " @ ")
}

// Clean 画布清屏
func (d *Drawer) Clean() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// 根据算法更新画布的刷新时间
func (d *Drawer) refreshInterval() {
	bodyLen := time.Duration(len(d.Snake.Body) * 20)
	d.Interval = d.Interval - bodyLen*time.Millisecond
	if d.Interval < 100*time.Millisecond {
		d.Interval = 100 * time.Millisecond
	}
}
