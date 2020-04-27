// in array which contains speed data for every second
// search 10-seconds period of time with maximum average speed
// speed data for array ar generated randomly
package main

import (
	"fmt"
	"math/rand"
	"time"
)

var data [100]int

const minPossibleSpeed = 10
const maxPossibleSpeed = 100
const sliceSize = 10

func main() {
	var slice []int
	var maxSpeedAverage float32 = 0
	var tempSpeedAverage float32
	var indexMaxSpeedAverage int = -1

	// initialize random generator with the same values each time
	// rand.Seed(42)

	// initialize random generator with different values each time
	rand.Seed(time.Now().UnixNano())

	// generate speed data for array randomly
	for i := range data {
		data[i] = minPossibleSpeed + rand.Intn(maxPossibleSpeed-minPossibleSpeed+1)
	}

	fmt.Println("Speed data : ", data) // this print is only for demo purpose

	// find 10-seconds period of time with maximum average speed
	for i := 0; i <= len(data)-sliceSize; i++ {
		slice = data[i : i+sliceSize]
		tempSpeedAverage = 0
		for j := 0; j < sliceSize; j++ {
			tempSpeedAverage += float32(slice[j])
		}
		tempSpeedAverage /= sliceSize
		if tempSpeedAverage > maxSpeedAverage {
			maxSpeedAverage = tempSpeedAverage
			indexMaxSpeedAverage = i
		}
	}
	fmt.Println("Maximum average speed =", maxSpeedAverage, "corresponds 10-seconds period of time starting from", indexMaxSpeedAverage, "second.")
}
