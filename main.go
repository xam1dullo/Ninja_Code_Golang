package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	//switch 
	switch opt {
	case "a":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price: ", reader)
		p,err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("the price must be a number")
			promptOptions(b)
		}
		b.addItem(name, p)
		fmt.Println("Item added - ", name,price)
		promptOptions(b)
	case "t":
		tip, _ := getInput("enter Tip amount ($): ", reader)
		t,err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("the tip must be a number")
			promptOptions(b)
		}
		b.upadateTip(t)
		fmt.Println("Tip added- ", tip)
		promptOptions(b)
	case "s":
		fmt.Println("you chose to save the bill",b) 
		 
	default:
		fmt.Println("Invalid option...")
	}
}

func main() {
	mybill := createBill()
	promptOptions(mybill)
	fmt.Println(mybill)
}
