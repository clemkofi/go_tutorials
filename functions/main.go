package main

import (
	"fmt"
	"math"
	"strings"
)

func sayGreeting(name string) {
	fmt.Printf("Good day %v \n", name)
}

func sayBye(name string) {
	fmt.Printf("Goodbye %v \n", name)
}

// passing functions as arguments to other functions
func cycleNames(names []string, f func(n string)) {
	for _, v := range names {
		f(v)
	}
}

// fucntion with return type
func circleArea(r float64) float64 {
	return math.Pi * r * r
}

// function with multiple return values
func getInitials(name string) (string, string) {
	caps := strings.ToUpper(name)
	names := strings.Split(caps, " ")

	var initials []string

	for _, v := range names {
		initials = append(initials, v[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], " "
}

func main() {
	// sayGreeting("Muriel")
	// sayGreeting("Naa")
	// sayBye("Mami")

	cycleNames([]string{"Muriel", "Naa", "Mami"}, sayGreeting)
	cycleNames([]string{"Muriel", "Naa", "Mami"}, sayBye)

	a1 := circleArea(20.5)

	fmt.Println(a1)
	fmt.Printf("the area of the circle is %0.3f \n", a1)

	fInitial1, lInitial1 := getInitials("Muriel Ofei-Quartey")
	fInitial2, lInitial2 := getInitials("Clement Nartey")
	fInitial3, lInitial3 := getInitials("Naa")

	fmt.Println(fInitial1, lInitial1)
	fmt.Println(fInitial2, lInitial2)
	fmt.Println(fInitial3, lInitial3)

}
