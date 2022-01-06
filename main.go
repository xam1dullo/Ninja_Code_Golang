package main

import (
	"fmt"
)

func main() {
	// var ages [3]int = [3]int{21, 13, 25}
	var ages = [3]int{21, 13, 25}
	names := [4]string{"xamidullo","azizbek","xumoyun","mr.xamidullo"}
	names[0] = "xamidullox"
	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	// slices (use arrays under the hood)
	var scores = []int{100,12,25,64}
	scores[3] =123
	scores = append(scores, 125)
	fmt.Println(scores,len(scores))


	//slice ranges

	rangeOne := names[1:4]
	rangeTwo := names[:3]
	rangeTree := names[3:]

	fmt.Println(rangeOne ,len(rangeOne))
	
	fmt.Println(rangeTwo, len(rangeTwo))
	
	fmt.Println(rangeTree, len(rangeTree))
}