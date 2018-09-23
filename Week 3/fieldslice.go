package main

import (
	"fmt"
)

type person struct {
	fname string
	age   int
	prefs []string
}

func main() {
	p1 := person{
		"Bond",
		32,
		[]string{"who", "ate", "the", "last", "piece"},
	}
	fmt.Println(p1.prefs)
}
