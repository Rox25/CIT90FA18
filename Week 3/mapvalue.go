package main

import (
	"fmt"
)

func main() {
	x := map[string]int{"Jack": 30, "Nathanial": 25, "Laura": 55}
	for k, _ := range x {
		fmt.Println(k, "-", v)
	}
}
