package main

import (
	"fmt"
)

type person struct{
	fName string
	lName string
	favFood []string
}

func main() {
	p1 := person{
		fName: "Louie",
		lName: "Poodle",
		favFood: []string{"ScoobySnacks", "beef jerky", "chicken"},
	}
	fmt.Println(p1)
	fmt.Println(p1.fName)
	fmt.Println(p1.favFood)

	for  _, v := range p1.favFood {
		fmt.Println(v)
	}
}