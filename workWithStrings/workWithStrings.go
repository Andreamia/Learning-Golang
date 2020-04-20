// analyze entered string for order of opening and closing symbols
// colorize opening and closing symbols in string
// find the longest string part in alphabetical order of english alphabet

package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/fatih/color"
)

// function analize(stroka, openSymbol, closeSymbol string)
// return string with results of opening and closing symbols order's analysis in stroka string
// analysis is performed by putting opening brackets into and out of stack array
// parameters:
//   stroka - string for analysis
//   openSymbol - opening symbol for analysis
//   closeSymbol - closing symbol for analysis

func analyze(stroka, openSymbol, closeSymbol string) {

	var stack []string // array for stack of opening symbols
	var index int = 0  // index for opening symbols in stack
	var result string  // result with analysis of opening and closing symbols order

	for _, char := range stroka {
		switch string(char) {
		case openSymbol:
			{
				stack = append(stack, string(char))
				index++
			}
		case closeSymbol:
			{
				if index == 0 || stack[index-1] != openSymbol {
					result = "INCORRECT"
					goto AnalysisResult
				} else {
					index--
				}
			}
		}

	}
	if index != 0 {
		result = "INCORRECT"
	} else {
		result = "CORRECT"
	}
AnalysisResult:
	fmt.Println("Order of", openSymbol, "and", closeSymbol, "in string is ", result, "!")
}

// colorize opening and closing symbols in the string
// corresponding opening and closing symbols have the same color
// parameters:
//   stroka - string for analysis
//   openSymbol - opening symbol for analysis
//   closeSymbol - closing symbol for analysis
func colorize(stroka, openSymbol, closeSymbol string) {

	var stack []string                         // array for stack of opening symbols
	var index int = 0                          // index for opening symbols in stack
	var colors []color.Attribute               // array with colors for opening and closing symbols
	var colorsClosingNonPair []color.Attribute // array with colors for non-pair closing symbols
	var count int = 0                          // index for non-pair closing symbols

	colors = append(colors, color.FgRed, color.FgBlue, color.FgYellow,
		color.FgCyan, color.FgMagenta, color.FgGreen)

	colorsClosingNonPair = append(colorsClosingNonPair, color.FgHiBlue, color.FgHiYellow,
		color.FgHiCyan, color.FgHiMagenta, color.FgHiGreen, color.FgHiRed)

	fmt.Print("Your string with colored opening and closing symbols: ")
	for _, char := range stroka {
		switch string(char) {
		case openSymbol:
			{
				stack = append(stack, string(char))
				color.Set(colors[index%6])
				fmt.Print(string(char))
				color.Unset()
				index++
			}
		case closeSymbol:
			{
				if index == 0 || stack[index-1] != openSymbol {
					color.Set(colorsClosingNonPair[count%6])
					fmt.Print(string(char))
					color.Unset()
					count++
				} else {
					color.Set(colors[(index-1)%6])
					fmt.Print(string(char))
					color.Unset()
					index--
				}
			}
		default:
			{
				fmt.Print(string(char))
			}
		}

	}
}

// find the longest string part in alphabetical order of english alphabet
// upper and lower case of characters is ignored
func findLongestAlphabeticalPart(stroka string) {
	var strArray []string
	var runeArray []rune
	var i int = 0
	var strTemp1, strTemp2 string = "", ""
	var temp1, temp2 rune = 0, 0

	for _, char := range stroka {
		runeArray = append(runeArray, char)
		strArray = append(strArray, string(char))
		temp1 = runeArray[i]
		if i != 0 {
			temp2 = runeArray[i-1]
		}
		if temp1 >= 97 && temp1 <= 122 {
			temp1 = temp1 - 32
		}
		if temp2 >= 97 && temp2 <= 122 {
			temp2 = temp2 - 32
		}
		if temp1 >= 65 && temp1 <= 90 {
			if i == 0 {
				strTemp1 = strArray[0]
			} else if temp1 > temp2 {
				strTemp1 += strArray[i]
			}
			if len(strTemp2) < len(strTemp1) {
				strTemp2 = strTemp1
			}
			if temp1 < temp2 {
				strTemp1 = strArray[i]
			}
		} else {
			strTemp1 = ""
		}
		i++
	}
	if len(strTemp2) < len(strTemp1) {
		strTemp2 = strTemp1
	}

	fmt.Println("The longest string part in alphabetical order is: ", strTemp2,
		". It's length = ", len(strTemp2))

}

// program execution
func main() {

	var str string     // string for analysis
	var opening string // opening symbol for analysis
	var closing string //closing symbol for analysis

	// read user's string from terminal
	fmt.Println("Enter string for analysis: ")
	reader := bufio.NewReader(os.Stdin)
	str, _ = reader.ReadString('\n')

	// read opening symbol from terminal
	fmt.Println("Enter opening symbol: ")
	opening, _ = reader.ReadString('\n')
	opening = opening[:1]

	// read closing symbol from terminal
	fmt.Println("Enter closing symbol: ")
	closing, _ = reader.ReadString('\n')
	closing = closing[:1]

	fmt.Println("---")

	// colorize opening and closing symbols in the string
	colorize(str, opening, closing)

	// analyze opening and closing symbols order in entered string
	analyze(str, opening, closing)

	// find the longest string part in alphabetical order of english alphabet
	// upper and lower case of characters is ignored
	findLongestAlphabeticalPart(str)

}
