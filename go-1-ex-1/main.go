package main

import "fmt"

func main() {
	// TODO: Declare and initialize the variables being used in the output!
	var firstName, lastName, dayOfBirth, monthOfBirth, yearOfBirth, numberOfSiblings, heightInMeters, zodiacSign = "Levin", "schmid", 28, 04, 9, 2, 2.90, '\u2649' 
	fmt.Printf("Vor- und Nachname: %s %s\n", firstName, lastName)
	fmt.Printf("Geburtsdatum: %d.%d.%d\n", dayOfBirth, monthOfBirth, yearOfBirth)
	fmt.Printf("Anzahl Geschwister: %d\n", numberOfSiblings)
	fmt.Printf("Grösse (in Metern): %.2f\n", heightInMeters)
	fmt.Printf("Sternzeichen: %c\n", zodiacSign)
}
