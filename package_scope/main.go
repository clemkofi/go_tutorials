package main

import "fmt"

var score = 60.2

func main() {
	// access function that is created in a different file but in the same package
	sayHello("Clement")

	for _, v := range points {
		fmt.Println(v)
	}

	showScore()
}
