package main

import (
	"fmt"
	"math/rand"
)

func main() {
	randomNumber := (rand.Intn(100))
	// fmt.Print(randomNumber) <-- prints pseudo-random integer

	// Create an if/else condition to check whether your number is greater or not than 50
	if randomNumber > 50 {
		fmt.Printf("my random number is %d and is greater than %d\n", randomNumber, 50)
	} else {
		return
	}

	// Store in the variable school the string Holberton School
	school := "Holberton School"
	// Create an if/else condition to check whether the string school is equal to Holberton School
	if school == "Holberton School" {
		fmt.Printf("I am a student of the %s\n", school)
	} else {
		return
	}

	// Store in the variable beautifulWeather the boolean value true
	beautifulWeather := true
	// Create an if/else condition to check whether the variable beautifulWeather is true or not
	if beautifulWeather {
		fmt.Printf("It's a beautiful weather!\n")
	} else {
		return
	}

	// Store in the variable holbertonFounders the following values Rudy Rigot, Sylvain Kalache, Julien Barbier
	holbertonFounders := []string{"Rudy Rigot", "Sylvain Kalache", "Julien Barbier"}
	//Create a for loop that will iterate over the holbertonFounders
	for i := 0; i < 3; i++ {
		fmt.Printf("%v is a founder\n", holbertonFounders[i])
	}
}
