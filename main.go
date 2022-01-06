package main

import (
	"fmt"
	"math"
)

func sayGreeting(n string) {
	fmt.Printf("Good morning %v\n", n)
}

func sayBae(n string) {
	fmt.Printf("Goodbye %v\n", n)
}

func cycleNames(n []string, f func(string)){
	for _, v := range n {
		f(v)
	}
}

func circleArea(r float64) float64{
	return math.Pi * r * r
}

func main() {
	// sayGreeting("xamidullo")
	// sayGreeting("azizbek")
	// sayBae("Humoyun")
	
	
	cycleNames([]string{"xamidullo", "azizbek", "humoyun"}, sayGreeting)
	cycleNames([]string{"xamidullo", "azizbek", "humoyun"}, sayBae)

	a1 := circleArea(10.5)
	a2 := circleArea(15)
	fmt.Println(a1, a2)
	fmt.Printf("Circule 1 is %0.3f and circle 2 is %0.3f \n",a1, a2)
}