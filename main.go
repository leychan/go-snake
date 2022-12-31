package main

import (
	"github.com/eiannone/keyboard"
	"time"
)

func main() {

	//初始化画布大小
	d := &Drawer{
		Height: 10,
		Width:  20,
	}

	//协程获取用户输入的方向
	go func(d *Drawer) {
		for {
			_, key, _ := keyboard.GetSingleKey()
			d.Snake.Direct = key
		}
	}(d)

	//画布初始化以及贪吃蛇初始化
	d.Init()

	//贪吃蛇持续运行,画布每隔一段时间就刷新,直到贪吃蛇死亡
	for {
		d.Clean()
		d.Run()
		time.Sleep(d.Interval)
	}

}
