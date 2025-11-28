package main

import "fmt"

type FullName struct {
	// TODO: add fields
	firstName string
	lastName  string
}

// TODO: declare a structure for birth date
type Birthdate struct{
	day		int
	month	int
	year	int
}

type Profile struct {
	// TODO: embed full name and birth date information
	birthdate 		 Birthdate
	fullName 		 FullName
	NumberOfSiblings byte
	ZodiacSign       rune
}

func main() {
	var me = Profile{
		// TODO: set name and birth date information
		birthdate: Birthdate{
			day: 1,
			month: 1,
			year: 2000,
		},
		fullName: FullName{
			firstName: "Levin",
			lastName: "Schmid",
		},
		NumberOfSiblings: 2,   // TODO: adjust
		ZodiacSign:       'S', // TODO: adjust
	}
	fmt.Println(me)

	fmt.Println("Siblings Before:", me.NumberOfSiblings)
	// TODO: imagine, you get a little brother or sister
	me.NumberOfSiblings = me.NumberOfSiblings+1
	fmt.Println("Siblings After:", me.NumberOfSiblings)
}
