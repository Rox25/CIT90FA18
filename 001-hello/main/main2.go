package main

import (
	"fmt"
)

func foo() {


name :="James" 
str := `html here` + ` ` + name + ` `+  `more html`
fmt.Println(str)

s := fmt.Sprintln(`mas ` + ` ` + name +` `+  `menos`)
fmt.Println(s)

}