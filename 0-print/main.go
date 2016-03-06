package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello, we are Holberton School")

	// Get the current time.
	t := time.Now()
	fmt.Printf("the date is %s\n", t.Local())

	//specify the year
	fmt.Println("the year is", t.Year())

	//specify the month
	fmt.Println("the month is", t.Month())

	//specify the day
	fmt.Println("the day is", t.Day())

}
