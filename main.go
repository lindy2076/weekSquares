package main

import (
	"flag"
	"log"
	"time"
)

// todo: customize colors of everything, symbols for everything, dob
func main() {
	years := 80
	var t2 string
	flag.StringVar(&t2, "dob", "2023-07-01", "your date of birth")
	flag.Parse()

	dob, err := time.Parse(time.DateOnly, t2)
	if err != nil {
		log.Fatal("The date passed is badly formated!!: ", err)
	}
	if dob.After(time.Now()) {
		log.Fatal("The date passed is after current date!!")
	}
	diff := time.Since(dob)

	weeksPassed := uint(diff.Hours() / 24 / 7)

	printSquares(uint(years), weeksPassed, false)
}
