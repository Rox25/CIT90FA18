package main

import (
	"fmt"
)

type person struct {
	last string
	age  int
}

func main() {
	p1 := person{"Goldie", 59}
	p2 := person{"Lancaster", 32}
	p1.foo()
	p2.foo()
}

func (p person) foo() {
	fmt.Println("Hi!", p.last)
}
