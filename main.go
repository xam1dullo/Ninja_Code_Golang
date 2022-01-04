package main

import "fmt"

func main() {
	age := 21
	name := "Xamidullo"
	
	//Print
	fmt.Print("Hello ")
	fmt.Print("Golang! \n")
	fmt.Print("new Line! \n")

	fmt.Println("Hello ninjas!")
	fmt.Println("Goodbye ninjas!")
	fmt.Println("My age is ", age, "and My name is ",name)

	//Print (formatted string) %_ = format specifer
	fmt.Printf("my age is %v and my name is %v \n",age, name)
	fmt.Printf("my age is %v and my name is %q \n",age, name)
	fmt.Printf("age is of type %T \n",age)
	fmt.Printf("you scored %f points! \n", 222.25)

	// fmt.Sprint(save formatted string)
	var str = fmt.Sprintf("my age is %v and my name is %v \n",age, name)

	fmt.Println("the saved string is :",str)
}