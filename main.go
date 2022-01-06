package main

import (
	"fmt"
	"strings"
)

func getIntials(n string) (string, string) {
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")
	
	var initials []string
	for _, v := range names {
		initials = append(initials, v[:1])
	}
	if len(initials) > 1 {
		return initials[0], initials[1]
	}
	return initials[0], "_"
}


func main() {
	fn1, sn1 :=	getIntials("tifa lochart")
	fmt.Println(fn1, sn1)
	fn2, sn2 :=	getIntials("cloud strinf")
	fmt.Println(fn2, sn2)
	fn3, sn3 :=	getIntials("xamidullox azizbek")
	fmt.Println(fn3, sn3)
	fn4, sn4 :=	getIntials("xamidullox")
	fmt.Println(fn4, sn4)
}