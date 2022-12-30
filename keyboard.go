package main

import (
	"github.com/eiannone/keyboard"
	"os"
)
import "fmt"

func GetInput() {

	for {
		input, key, err := keyboard.GetSingleKey()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if key == keyboard.KeyCtrlQ {
			os.Exit(1)
		}

		fmt.Println(input)
	}

}
