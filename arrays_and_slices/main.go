package main

import "fmt"

func main() {
	// arrays == they always have a fixed length when declared in Go
	var ages [3]int = [3]int{20, 40, 50}
	// var agesTwo = [3]int{20, 40, 50}

	names := [4]string{"Muriel", "Naa", "Lamiokor", "Mami"}

	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	// slices == they work like the dynamic arrays from other languages == use arrays under the hood
	scores := []int{100, 50, 40}
	scores[2] = 45
	scores = append(scores, 70) // used to append a value to a slice

	fmt.Println(scores, len(scores))

	// slice ranges == used when we want to get slices from arrays or other slices
	rangeOne := names[1:3]
	rangeTwo := names[2:]

	fmt.Println(rangeOne, rangeTwo)

	// slice ranges can be appended since they are deep copies of the actual values
	rangeOne = append(rangeOne, "Ofei-Quartey")
	fmt.Println(rangeOne)
}
