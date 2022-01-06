package main

import "fmt"

func main() {
	menu := map[string]float64{
		"soup":               1.50,
		"price":              2.50,
		"salad":              3.00,
		"dessert":            4.00,
		"toffee pudding pie": 5.00,
	}
	fmt.Println(menu)
	fmt.Println(menu["price"])


	// looping maps

	for key, value := range menu{
		fmt.Println(key, "- ",value)
	}
	// ints as key type/

	phonebook := map[int]string{
		1: "James",
		2: "John",
		3: "Robert",
		4: "Michael",
		5: "William",
		6: "David",
		7: "Richard",
		8: "Joseph",
	}
	fmt.Println(phonebook)
	fmt.Println(phonebook[3])
	phonebook[1] = "Jill"
	fmt.Println(phonebook[5])
	phonebook[5] = "Jack"
	fmt.Println(phonebook)
}