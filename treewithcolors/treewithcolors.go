// draw a multi-colored decorated Christmas tree by symbols in terminal
// install "github.com/fatih/color" package before run: go get github.com/fatih/color

package main

import (
	"fmt"

	"github.com/fatih/color"
)

// print character with colour
func printColoredCharacter(character string, colour color.Attribute) {
	color.Set(colour)
	fmt.Print(character)
	color.Unset()
}

//program execution
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
						printColoredCharacter("|", color.FgMagenta)
					} else {
						if w%5 == 0 {
							printColoredCharacter("+", color.FgRed)
						} else {
							printColoredCharacter("*", color.FgGreen)
						}
					}
				}
			}
		}
		fmt.Println()
	}

}
