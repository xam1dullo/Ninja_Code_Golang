package main

import (
	"fmt"
)

func main() {
	mybill := newBill("mario's bill")

	mybill.addItem("pie", 5.88)
	mybill.addItem("cake", 1.88)
	mybill.addItem("coffe", 2.77)
	mybill.addItem("toffee", 4.95)
	mybill.addItem("onion soup", 4.53)
	mybill.upadateTip(10)

	fmt.Println(mybill.format())
}
