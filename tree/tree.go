package main

import "fmt"

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
                        fmt.Print("|")
                    } else {
                        if j%5 == 0 {
                            fmt.Print("+")
                        } else {
                            fmt.Print("*")
                        }
                    }
                }
            }
        }
        fmt.Println()
    }

}