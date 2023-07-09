package main

import (
	"log"
	"time"
)

// todo: customize colors of everything, symbols for everything, dob
func main() {
	years := 80
	t2 := "2003-11-17"

	dob, err := time.Parse(time.DateOnly, t2)
	if err != nil {
		log.Fatal("The date passed is badly formated!!", err)
	}
	diff := time.Since(dob)

	weeksPassed := uint(diff.Hours() / 24 / 7)

	printSquares(uint(years), weeksPassed, false)
}
