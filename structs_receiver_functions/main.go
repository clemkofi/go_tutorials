package main

import "fmt"

func main() {
	testBill := newBill("Naa's Bill")

	testBill.addTip(20)

	testBill.addItem("Fufu", 20.56)
	testBill.addItem("Fried Rice", 15.99)
	testBill.addItem("Yam", 8.99)

	fmt.Println(testBill.format())
}
