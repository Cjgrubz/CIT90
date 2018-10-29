package main

import (
	"fmt"
)

func main() {
	m := [string]int{"james": 7, "jenny": 8}
	fmt.Printf("%T\n", m)
	fmt.Println(m)

	fmt.Println(m["James"])

}
/*add*/
