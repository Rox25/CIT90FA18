package main

import (
	"fmt"
)

func main() {
	fmt.Println("Begin")
foo("James Bond")
foo("Miss MoneyPenny")
	fmt.Println("about to end")

}

// foo takes a value TYPE string
// we could also say:
//foo has a peramiter which is a VALUE of TYPE string
func foo(name string) {
	fmt.Println("Hello" , name)
}