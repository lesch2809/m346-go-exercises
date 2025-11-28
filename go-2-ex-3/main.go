package main

import "fmt"

func main() {
	// TODO: create a map called "modules"
	modules := map[uint]string{
		104: "das eine modul",
		117: "das andere Modul",
		346: "das letze modul",
	}

	fmt.Println("Modul 104:", modules[104])
	fmt.Println("Modul 117:", modules[117])
	fmt.Println("Modul 346:", modules[346])

	// TODO: delete 
	delete(modules, 104)
	// TODO: add one
	modules[187] = "uk modul"
	// TODO: replace one
	modules[117] = "das neue modul"
	fmt.Println(modules)
}
