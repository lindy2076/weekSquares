package main

import (
	"flag"
	"log"
	"time"
)

// todo: customize colors, symbols of everything, dob
func main() {
	var yearsExpected uint
	var userDoB string
	flag.StringVar(&userDoB, "dob", "2023-07-09", "your date of birth")
	flag.UintVar(&yearsExpected, "y", 80, "your life expectancy")
	flag.Parse()

	dob := parseDoB(userDoB)
	sinceDoB := time.Since(dob)

	weeksPassed := uint(sinceDoB.Hours() / 24 / 7)

	printSquares(uint(yearsExpected), weeksPassed, false)
}

func parseDoB(userflag string) time.Time {
	dob, err := time.Parse(time.DateOnly, userflag)
	if err != nil {
		log.Fatal("The given date is badly formatted")
	}
	if dob.After(time.Now()) {
		log.Fatal("The given date is after")
	}
	return dob
}
