package main

import (
	"fmt"
)

func main() {
	// string
	var nameOne string = "Xamidullo"
	var nameTwo = "Azizbek"
	var  nameTree string

	fmt.Println(nameOne, nameTwo)

	nameOne = "Peach"
	nameTree = "Browser"

	fmt.Println(nameOne, nameTwo, nameTree)

	nameTree = "Yoshi"

	fmt.Println(nameTree)

	//ints

	var ageOne = 21
	var ageTwo = 12
	ageTree := 25;

	fmt.Println(ageOne, ageTwo, ageTree)

	//bits & memory

	var numOne int8 = 25
	var numTwo int8 = -128
	var numTree uint16 = 255
	
	fmt.Println(numOne, numTwo, numTree);

	//floats

	var scoreOne float32 = 25.45
	var scoreTwo float32 = 88888888888888888888888888888888888888.111
	scoreTree := 1.5
	fmt.Println(scoreOne, scoreTwo, scoreTree)
}