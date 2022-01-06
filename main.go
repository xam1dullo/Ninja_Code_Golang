package main

import (
	"fmt"
	"sort"
)

func main() {
	// greeting := "hello there friends!"
	//string methods
	// fmt.Println(strings.Contains(greeting,"hello "))
	// fmt.Println(strings.ReplaceAll(greeting,"hello","hi"))
	// fmt.Println((strings.ToUpper(greeting)))
	// fmt.Println(strings.Index(greeting,"ll"))
	// fmt.Println(strings.Index(greeting,`ll`))
	// fmt.Println(strings.Split(greeting, " "))
	
	//sort methods
	ages  := []int{23,12,42,553,64,757,8,68,6,4,2}
	
	sort.Ints(ages)
	fmt.Println(ages)
	
	index := sort.SearchInts(ages,30)
	fmt.Println(index) 
	
	names := []string{"xamidullo","azizbek","humoyun","xamidullo","azizbek","humoyun"}
	sort.Strings(names)
	fmt.Println(names)
	fmt.Println(sort.SearchStrings(names,"xamidullo"))
} 
	
	