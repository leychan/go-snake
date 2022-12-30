package main

import (
	"github.com/eiannone/keyboard"
	"time"
)

func main() {
	//var ch chan int
	d := &Drawer{
		Length: 10,
	}

	go func(d *Drawer) {
		for {
			_, key, _ := keyboard.GetSingleKey()
			d.Direct = key
		}
	}(d)

	d.Init()
	d.InitHeader()
	d.RandomDirect()
	d.newFood()
	d.Print()
	for {
		d.Clean()
		d.MoveDirect()
		d.Print()
		time.Sleep(time.Second)
	}

}
