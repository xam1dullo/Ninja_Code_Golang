package main

import "fmt"

func main() {
	// while loop
	x := 0
	for x < 5 {
		fmt.Println("value of x is ", x)
		x++
	}
	
	//for loop
	for i :=0; i < 5; i++{
		fmt.Println("value of i is ", i)
	}
	
	names := []string{"xamidullo","azizbek","humoyun","hojiakbar","begzodbek"}

	for i :=0; i < len(names); i++{
		fmt.Println("value of names[",i,"] is ", names[i])
	}
	
	// for in loop
	
	for index, value := range names{
		fmt.Printf("the position at index %v is %v \n",index,value)
	}

	for _, value := range names{
		fmt.Printf("the value is %v \n",value)
		value = "new String"
		fmt.Println(names)
	}
	fmt.Println(names)
}