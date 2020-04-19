// draw a multi-colored decorated Christmas tree by symbols in terminal
// install "github.com/fatih/color" package before run: go get github.com/fatih/color

package main

import (
	"fmt"

	"github.com/fatih/color"
)

func main() {

	var height int = 20          // tree height
	var width int = height*2 + 1 // tree width

	for h := 0; h < height; h++ {
		for w := 0; w < width; w++ {
			if w < (width/2 - h) {
				fmt.Print(" ")
			} else {
				if w > (width/2 + h) {
					fmt.Print(" ")
				} else {
					if w == width/2-1 || w == width/2 || w == width/2+1 {
						color.Set(color.FgMagenta)
						fmt.Print("|")
						color.Unset()
					} else {
						if w%5 == 0 {
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
