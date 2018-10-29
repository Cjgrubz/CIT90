package main 

import (
	"fmt"
)

func main () {
	x := 42
	y := 43

	xi := []int{x, y}

	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", xi)
	fmt.Println(xi)
}
/*add*/