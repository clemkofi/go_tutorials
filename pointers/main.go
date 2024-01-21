package main

import "fmt"

// if you want to pass a pointer to a function you do so by adding * to the datatype
func updateName(n *string) {
	*n = "Clement"
}

func main() {
	name := "Muriel"

	m := &name // the & allows us to get the memory reference of the variable

	fmt.Println("the name is:", name)

	fmt.Println("the memory address of name is:", m)

	updateName(m)

	fmt.Println("the updated name is:", name)

}
