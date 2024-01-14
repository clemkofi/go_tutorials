package main

import "fmt"

func main() {
	// fmt.Println("Hello World!")

	// // strings
	// var nameOne string = "Muriel"
	var nameTwo = "Naa"
	// var nameThree string

	// fmt.Println(nameOne, nameTwo, nameThree)

	// nameOne = "Mami"
	// nameThree = "Lamiokor"
	// fmt.Println(nameOne, nameTwo, nameThree)

	// // short hand for declaring variables (can only be used within functions not outside)
	// nameFour := "Chantel"
	// fmt.Println(nameOne, nameTwo, nameThree, nameFour)

	// ints
	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40

	fmt.Println(ageOne, ageTwo, ageThree)

	// bits and memory
	// signed ints (+ and - numbers)
	var numOne int8 = 127
	var numTwo int8 = 25

	// unsigned ints (only positive numbers)
	var numThree uint8 = 255

	fmt.Println(numOne, numTwo, numThree)

	// floats (default is float64)
	var scoreOne float32 = 23.776
	var scoreTwo float64 = 3455692328434.9987

	scoreThree := 38.998

	fmt.Println(scoreOne, scoreTwo, scoreThree)

	// printing functions

	// print = everything is printed on the same line
	fmt.Print("Hello, ")
	fmt.Print("World! \n")

	// Println = adds a new line at the end of every print statement
	fmt.Println("Good night!")
	fmt.Println("My name is", nameTwo, "and my age is ", ageTwo)

	// Printf (formatted strings) used with format specifiers
	fmt.Printf("My name is %v and the age is %v \n", nameTwo, ageTwo)
	fmt.Printf("Her name is %q \n", nameTwo)                       // this puts quotes around strings and only used for strings
	fmt.Printf("age is of type %T \n", ageThree)                   //this gives us the type of variable used
	fmt.Printf("the total price is %f \n", 222.563)                // used to print floats but prints up to a precision of about 5 or 6
	fmt.Printf("ths formatted total price is %0.2f \n", 3324.8854) // this will print format to two decimal places

	// Sprintf = save the formatted string into a variable
	strSaved := fmt.Sprintf("My name is %v and the age is %v \n", nameTwo, ageTwo)
	fmt.Println("The saved string is:", strSaved)
}
