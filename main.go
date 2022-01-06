package main

import "fmt"

func updateName(x *string) {
	*x = "azizbek"
}

func main() {
	// gruop A = type  -> string, int, float, bool, arrays, structs
	name := "xamidullo"
	
	// updateName(name)

	fmt.Println("memory address of name: ", &name)
	m := &name
	fmt.Println("memory address : ", m)
	fmt.Println("value of namae: ", *m)
	updateName(m)
	fmt.Println("value of name: ", *m)

	// gruop B = types => slices, maps, functions
/* 
|-----name----|------m------|
| 0xc0000a0a0 | 0xc0000a0a0 |
| azizbek     | xamidullo   |
|-------------|-------------|

*/
	
}