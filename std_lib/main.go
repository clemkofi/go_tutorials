package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	greeting := "Hello there guys lo!"

	// none of the functions in the strings package implicitly reply the string  that is passe to it
	fmt.Println(strings.Contains(greeting, "Hello"))
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.ReplaceAll(greeting, "Hello", "Hi"))
	fmt.Println(strings.Index(greeting, "lo")) //gets the first index of the search value

	// the sort package == mostly implicitly modifies variable passed to it
	ages := []int{20, 70, 40, 32, 54, 99, 43, 22}

	sort.Ints(ages)
	fmt.Println(ages)
	fmt.Println(sort.SearchInts(ages, 54)) // can only search through slices that have been sorted with the sort package

	names := []string{"Muriel", "Naa", "Lamiokor", "Mami"}

	sort.Strings(names)
	fmt.Println(names)
	fmt.Println(sort.SearchStrings(names, "Naa"))

	////// if value for the search function does not exist it gives a value which is the same as the length of the slice provided + 1

}
