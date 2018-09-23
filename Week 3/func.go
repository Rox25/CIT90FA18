package main

import (
	"fmt"
)

type person struct {
	fName string
	lName string
}

func main() {
	// composite literal; struct literal
	p1 := person{"Roxanne", "Dutson"}
	p2 := person{"Jonah", "Smith"}
	fmt.Println(p1)
	fmt.Println(p2)
	n := meaning()
	fmt.Println(n)
}

// func (receiver) identifier(parameters) (returns) { <code> }

func meaning() int {
