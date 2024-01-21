package main

import "fmt"

func updateMenu(m map[string]float64) {
	m["kelewele"] = 4.99
}

func main() {
	menu := map[string]float64{
		"fufu":       20.33,
		"fried rice": 19.54,
		"yam":        14.88,
	}

	fmt.Println(menu)
	fmt.Println(menu["fufu"]) // get a single value from the map

	// loop through the map
	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	phonebook := map[int]string{
		243250391: "mum",
		543022179: "clement",
		242716532: "dad",
	}

	fmt.Println(phonebook)

	// update the value of a key in a map
	phonebook[543022179] = "kofi"
	fmt.Println(phonebook)

	// GO has pointer wrapper types that are referenced by pointers so whenever they are passed into a function
	// they are not copied by rather referenced from the original pointer.
	// the types that behave like this are -> maps, slices, functions
	// example
	updateMenu(menu) // when this run it adds a new entry into the map and mutates it
	fmt.Println(menu)

	// this does not work for primitive types like -> float, int, strings, arrays, structs

}
