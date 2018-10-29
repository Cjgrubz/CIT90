package main

import(
	"fmt"
)

func main() {
	xi := []int{2,3,4,5,6}

	for i := range xi {
		fmt.Println(i)
	}

	for i, v := range xi {
		fmt.Println(i, v)
	}
}