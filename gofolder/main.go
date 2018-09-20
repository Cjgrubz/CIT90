package main

import (
	"fmt"
)

func main() {
	foo("Cyd", 47)
	foo("garret", 23)
	fmt.Scanln()
}

func foo(x string, y int) {
	fmt.Println(x, y)
}