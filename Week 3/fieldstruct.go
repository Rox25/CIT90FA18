package main

import (
	"fmt"
)

type person struct {
	fname string
	age   int
}

func main() {
	p1 := person{
		"Joseph",
		55,
	}
	fmt.Println(p1.fname)
}
