package main

import "fmt"

func main() {
	// TODO: declare a type for Student (with first and last name)
	type student struct {
		firstname string
		lastname  string
	}
	// TODO: declare a type for Class (consisting of multiple students)
	type class struct{
		students []student
	}
	// TODO: declare a map of modules being attended by multiple classes
	modules := map[string][]class{
		"114": []class{
			{
			students: []student{
				{firstname: "Max", lastname: "Mustermann"},
				{firstname: "Erika", lastname: "Musterfrau"},
				{firstname: "Julia", lastname: "Mayer"},
			},
			},
			{
			students: []student{
				{firstname: "Tom", lastname: "Hoffmann"},
				{firstname: "Eva", lastname: "Schulz"},
				{firstname: "Mia", lastname: "Klein"},
			},
		},
		
		},
		"117": []class{
			{
			students: []student{
				{firstname: "Hans", lastname: "Schmidt"},
				{firstname: "Anna", lastname: "Schneider"},
				{firstname: "Lena", lastname: "Fischer"},
			},
		},
			
		},
		"120": []class{
			{
			students: []student{
				{firstname: "Peter", lastname: "Fischer"},
				{firstname: "Laura", lastname: "Weber"},
				{firstname: "Sophie", lastname: "Becker"},
			},
		},
		},
	}
	// @Patrick Bucher: Wie kann ich meherer klassen zu eine kurs hinzufuegen?
	// TODO: output everything using fmt.Println()
	fmt.Println(modules)
}
