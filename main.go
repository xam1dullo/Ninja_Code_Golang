package main

import "fmt"

func main() {
	age := 45
	fmt.Println(age <= 50)
	fmt.Println(age >= 50)
	fmt.Println(age == 50)
	fmt.Println(age != 50)
	
	if age < 30{
		fmt.Println("age is less then 30")
		}else if age < 40 {
			fmt.Println("age is less then 40")
			}else {
				fmt.Println("age is less then 45")
			}
			
			names := []string{"John", "Paul", "George", "Ringo"}
			for index, value := range names {
				if index == 1{
					fmt.Println("continuning at pos", index)
					continue
				}
				if index == 2{
					fmt.Println("breaking at pos", index)
					break
				}
				fmt.Println("the value at", index, "is", value)
			}
		}