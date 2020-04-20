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

	var counter int = 0 // counter for opening symbols from stroka string
	var result string   // result with analysis of opening and closing symbols order

	for _, char := range stroka {
		switch string(char) {
		case openSymbol:
			{
				counter++
			}
		case closeSymbol:
			{
				if counter != 0 {
					counter--
				} else {
					result = "INCORRECT"
					goto AnalysisResult
				}
			}
		}

	}
	if counter == 0 {
		result = "CORRECT"
	} else {
		result = "INCORRECT"
	}
AnalysisResult:
	fmt.Println("Order of", openSymbol, "and", closeSymbol, "in string is", result, "!")
}

// print symbol with character ANCII code with colour
func printColoredCharacter(character rune, colour color.Attribute) {
	color.Set(colour)
	fmt.Print(string(character))
	color.Unset()
}

// colorize opening and closing symbols in the string
// corresponding opening and closing symbols have the same color
// parameters:
//   stroka - string for analysis
//   openSymbol - opening symbol for analysis
//   closeSymbol - closing symbol for analysis
func colorize(stroka, openSymbol, closeSymbol string) {

	var counter int = 0 // counter for opening symbols from stroka string
	var index int = 0   // index for non-pair closing symbols

	var colors []color.Attribute // array with colors for opening and closing symbols
	colors = append(colors, color.FgRed, color.FgBlue, color.FgYellow,
		color.FgCyan, color.FgMagenta, color.FgGreen)

	var colorsClosingNonPair []color.Attribute // array with colors for non-pair closing symbols
	colorsClosingNonPair = append(colorsClosingNonPair, color.FgHiBlue, color.FgHiYellow,
		color.FgHiCyan, color.FgHiMagenta, color.FgHiGreen, color.FgHiRed)

	fmt.Print("Your string with colored opening and closing symbols: ")
	for _, char := range stroka {
		switch string(char) {
		case openSymbol:
			{
				printColoredCharacter(char, colors[counter%6])
				counter++
			}
		case closeSymbol:
			{
				if counter != 0 {
					printColoredCharacter(char, colors[(counter-1)%6])
					counter--
				} else {
					printColoredCharacter(char, colorsClosingNonPair[index%6])
					index++
				}
			}
		default:
			{
				fmt.Print(string(char))
			}
		}

	}
}

// check is character ANCII code corresponds to ANCII code of upper case english letter or not
func isUpperCase(character rune) bool {
	return character >= 65 && character <= 90
}

// check is character ANCII code corresponds to ANCII code of lower case english letter or not
func isLowerCase(character rune) bool {
	return character >= 97 && character <= 122
}

// return ANCII code of upper case english letter corresponding to
// ANCII code of character lower case english letter
func changedToUpperCase(character rune) rune {
	return character - 32
}

// find the longest string part in alphabetical order of english alphabet
// upper and lower case of characters is ignored
func findLongestAlphabeticalPart(stroka string) {

	var index int = 0                      // index for stroka characters
	var strArray []string                  // array with characters from stroka string
	var runeArray []rune                   // array with ANCII codes of characters from stroka string
	var strTemp1, strTemp2 string = "", "" // variables for composing and finding the longest substring in alphabetical order
	var temp1, temp2 rune = 0, 0           // variables for comparing ANCII codes from runeArray

	for _, char := range stroka {
		runeArray = append(runeArray, char)
		strArray = append(strArray, string(char))
		temp1 = runeArray[index]
		if index != 0 {
			temp2 = runeArray[index-1]
		}
		if isLowerCase(temp1) {
			temp1 = changedToUpperCase(temp1)
		}
		if isLowerCase(temp2) {
			temp2 = changedToUpperCase(temp2)
		}
		if isUpperCase(temp1) {
			if index == 0 {
				strTemp1 = strArray[0]
			} else if temp1 > temp2 {
				strTemp1 += strArray[index]
			}
			if len(strTemp2) < len(strTemp1) {
				strTemp2 = strTemp1
			}
			if temp1 < temp2 {
				strTemp1 = strArray[index]
			}
		} else {
			strTemp1 = ""
		}
		index++
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
