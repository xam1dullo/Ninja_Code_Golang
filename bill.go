package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bills
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{"pie": 5.990, "cake": 2.999},
		tip:   0,
	}
	return b
}

// format the bill
func (b bill) format() string {
	fs := "Bill breakdown: \n"
	var total float64 = 0
	//list all items
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v \n", k+":", v)
		total += v
	}
	//total
	fs += fmt.Sprintf("%-25v ...$%0.2f", "totoal:", total )
	return fs
} 
