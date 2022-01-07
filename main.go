package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
func getInput(prompt string, r *bufio.Reader)(string, error){
	fmt.Print(prompt)
	input , err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}

func createBill() bill {
	reader :=bufio.NewReader(os.Stdin)
	
	name, _ := getInput("Create a new bill name: ", reader)
	b := newBill(name)	
	fmt.Println("Create the bill -", b.name)
	return b
}
func promptOptions (b bill){
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("Choose option (a -add item, s - save bill, s -add tip): ", reader)
	fmt.Println("Option: ", opt)
	switch opt {
	case "a":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price: ", reader)
		fmt.Println(name, price)
	case "t":
		tip, _ := getInput("enter Tip amount ($): ", reader)
		fmt.Println(tip)
	case "s":
		fmt.Println("Add tip")
	default:
		fmt.Println("Invalid option...")
	}
}

func main() {
	mybill := createBill()
	promptOptions(mybill)
	fmt.Println(mybill)
}
