package main

import (
	"fmt"

	"github.com/fatih/color"
)

func main() {

	var height int = 20
	var width int = height*2 + 1
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if j < (width/2 - i) {
				fmt.Print(" ")
			} else {
				if j > (width/2 + i) {
					fmt.Print(" ")
				} else {
					if j == width/2-1 || j == width/2 || j == width/2+1 {
						color.Set(color.FgMagenta)
						fmt.Print("|")
						color.Unset()
					} else {
						if j%5 == 0 {
							color.Set(color.FgRed)
							fmt.Print("+")
							color.Unset()

						} else {
							color.Set(color.FgGreen)
							fmt.Print("*")
							color.Unset()

						}
					}
				}
			}
		}
		fmt.Println()
	}

}
