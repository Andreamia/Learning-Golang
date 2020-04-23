// find prime numbers from 1 to number
// find nearest power of two for number

package main

import (
	"fmt"
	"math"
	"strconv"
)

// check if num is prime number
func isPrimeNumber(num int) bool {

	var half = int(num / 2)
	var reminder = 1
	var divider = 2

	for divider <= half && reminder != 0 {
		reminder = num % divider
		divider++
	}

	if reminder == 0 {
		return false
	} else {
		return true
	}
}

// print list of prime numbers from 1 to num
func printPrimeNumbers(num int) {

	fmt.Println("List of prime numbers from 1 to", num, ":")
	if num == 1 {
		fmt.Print("no prime numbers")
	}
	for i := 2; i <= num; i++ {
		if isPrimeNumber(i) {
			fmt.Print(i, " ")
		}
	}
	fmt.Println("")
}

// find nearest power of two for num
func printNearestPowerOfTwo(num float64) {

	var power, distance1, distance2 float64

	power = math.Floor(math.Log2(num))
	distance1 = math.Abs(math.Pow(2, power) - num)
	distance2 = math.Abs(math.Pow(2, power+1) - num)

	if distance1 == 0 || distance2 == 0 {
		fmt.Println("Nearest to " + strconv.Itoa(int(num)) + " power of 2 is = " + strconv.Itoa(int(num)) + " (=2^" + strconv.Itoa(int(math.Log2(num))) + ")")
	} else if distance1 < distance2 {
		fmt.Println("Nearest to " + strconv.Itoa(int(num)) + " power of 2 is = " + strconv.Itoa(int(math.Pow(2, power))) + " (=2^" + strconv.Itoa(int(power)) + ")")
	} else if int(distance1) == int(distance2) {
		fmt.Println("Nearest to " + strconv.Itoa(int(num)) + " power of 2 are = " + strconv.Itoa(int(math.Pow(2, power))) + " (=2^" + strconv.Itoa(int(power)) + ") and " + strconv.Itoa(int(math.Pow(2, power+1))) + " (=2^" + strconv.Itoa(int(power+1)) + ")")
	} else {
		fmt.Println("Nearest to " + strconv.Itoa(int(num)) + " power of 2 is = " + strconv.Itoa(int(math.Pow(2, power+1))) + " (=2^" + strconv.Itoa(int(power+1)) + ")")
	}
}

// execute program
func main() {
	var number int // number for analysis
	// read a number from terminal
	fmt.Println("Enter a whole number from 1 to 1000:")
	fmt.Scan(&number)

	if number >= 1 && number <= 1000 {
		fmt.Println("---")

		// find and print list of prime numbers from 1 to number
		printPrimeNumbers(number)

		// find and print nearest power of two for number
		printNearestPowerOfTwo(float64(number))

	} else {
		fmt.Println("You entered wrong value!")
	}

}
