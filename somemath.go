package main

import (
	"fmt"
)

const colorRed = "\033[0;91m"
const colorGreen = "\033[0;92m"
const colorYellow = "\033[0;93m"
const colorBlue = "\033[0;94m"
const colorNone = "\033[0m"
const fontBold = "\033[1m"

// TODO generate based on years
func stringFrom(yearsTotal uint, weeksPassed uint) string {
	// weeks := yearsTotal * 365 / 7

	rows, columns := 54, 78 // FIXME generate based on weeks
	res := " "

	for i := 0; i < rows*columns; i++ {
		if i < int(weeksPassed) {
			res += colorRed + "x" + colorNone
		} else if i == int(weeksPassed) {
			res += colorYellow + fontBold + "?" + colorNone
		} else {
			res += fontBold + "." + colorNone
		}
		if (i+1)%columns == 0 {
			res += "\n "
		}
	}
	return res
}

func printSquares(years uint, weeksPassed uint, trySquare bool) {
	fmt.Println(stringFrom(uint(years), weeksPassed))

	fmt.Printf("\nThis is your life. Every dot represents a week till you are %d yo.\n", years)
	fmt.Printf("Every 'x' is a week that can never be brought back. There are already %d 'x'.\n", weeksPassed)
	fmt.Println("Is there something valuable that you did this week..?")
}
