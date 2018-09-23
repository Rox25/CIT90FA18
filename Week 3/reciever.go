package main

import (
	"fmt"
)

type person struct {
	last string
	age  int
}

func main() {
	p1 := person{"Martha", 24}
	fmt.Println(p1)
}
