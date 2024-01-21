package main

import "fmt"

// to create a new type we uset the keyword type
type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// a function to be some sort of constructor to create a new bills
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}

	return b
}

// receiver functions are functions that can be associated to structs and objects
// they are very similar to methods of a class and can only be used with the objects they are associated to but cannot be used independently

// a receiver function to format the bill
// it is better to pass structs as pointers into receiver functions so that new copies are not made every single time
// structs passed as pointers in receiver functions are automatically dereferenced by GO
func (b *bill) format() string {
	fs := "Bill Breakdown \n\n"
	var total float64 = 0

	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v \n", k+":", v) // the -25 shows that you want the variable to be 25 characters long but you want the white space to be on the right of the text
		// if 25 would used then the white space would be on the left of the text instead

		total += v
	}

	// show tip
	fs += fmt.Sprintf("\n%-25v ...$%v", "tip:", b.tip)

	// show total
	fs += fmt.Sprintf("\n%-25v ...$%v", "total:", total)

	return fs
}

// receiver function to add a tip
func (b *bill) addTip(tip float64) {
	b.tip = tip
}

// receiver function to add a new item to the bill
func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}
