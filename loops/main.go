package main

import "fmt"

func main() {
	// // for loop 1st style
	// x := 0

	// for x < 5 {
	// 	fmt.Println("The value of x is:", x)
	// 	x++
	// }

	// // for loop 2nd style
	// for i := 0; i < 5; i++ {
	// 	fmt.Println("The value of i is:", i)
	// }

	names := []string{"Muriel", "Naa", "Lamiokor", "Mami"}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	// for loop 3rd style == similar to for in in python
	for index, value := range names {
		fmt.Printf("The value at index %v is %v \n", index, value)
	}

	// if you only need the value you can put an _ in place of the index
	for _, value := range names {
		fmt.Printf("The value of the name is %v \n", value)
	}

	/////// Conditionals

	// if - else
	age := 45

	fmt.Println(age <= 60)
	fmt.Println(age >= 60)
	fmt.Println(age == 45)
	fmt.Println(age != 50)

	if age < 30 {
		fmt.Println("the age is less than 30")
	} else if age < 40 {
		fmt.Println("the age is less than 40")
	} else {
		fmt.Println("the age is greater than 40")
	}

	for index, value := range names {
		if index == 1 {
			fmt.Println("continuing at index", index)
			continue
		}

		if index > 2 {
			fmt.Println("breaking at index ", index)
			break
		}

		fmt.Printf("The value at index %v is %v \n", index, value)
	}

}
