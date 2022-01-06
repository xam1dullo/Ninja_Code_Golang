package main

import "fmt"

func updateName(x string)string {
	x = "shubham"
	return x
}
func updateMenu(y map[string] float64){
	y ["coffe"] = 10000
}
func main() {
	// gruop A = type  -> string, int, float, bool, arrays, structs
	name := "xamidullo"
	name = 	updateName(name)
	fmt.Println(name)
	
	// gruop B = types => slices, maps, functions
	

	menu := map[string] float64{
		"coffe": 4.5,
		"tea": 2.5,
		"juice": 3.5,
	}
	
	updateMenu(menu)
	fmt.Println(menu)

}