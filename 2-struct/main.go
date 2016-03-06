package main

import (
	"fmt"
	"time"
)

//Create a user struct that has the following properties
type user struct {
	Name string
	DOB  string
	City string
}

func main() {
	//initialize the struct user in a variable called u with the following values
	u := user{Name: "Betty Holberton", DOB: "March 7, 1917", City: "Philadelphia"}

	//Create a method on the struct that will print Hello <Name>
	fmt.Printf("Hello %s\n", u.Name)

	t := time.Now()
	//fmt.Println(t.Year())

	value := "03/07/1917"
	layout := "01/02/2006"
	f, _ := time.Parse(layout, value)
	//fmt.Println(f.Year())

	integer := (t.Year() - f.Year())

	//Create a method on the struct that will produce the following input.  <Name> who was born in <City> would be XX years old today
	fmt.Printf("%s who was born in %s would be %d years old today.\n", u.Name, u.City, integer)
}
