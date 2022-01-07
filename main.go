package main

import (
	"fmt"
)

func main() {
	myBill := newBill("")
	
	fmt.Println(myBill.format())
	
}
