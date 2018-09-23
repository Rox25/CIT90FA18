package main

import (
	"fmt"
)

type person struct {
	fName string
	lName string
}

func main() {
	// composite literal
	p1 := person{"Harley", "Quinn"}
	p2 := person{"Hannah", "Stuart"}
	fmt.Println(p1)
	fmt.Println(p2)
}
