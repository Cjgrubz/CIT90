/*
Take the code from the previous exercise, then store the values of type person
in a map with the key of last name. Access each value in the map. Print out
the values, ranging over the slice.
*/
package main

import "fmt"

// Person : a representation of a person
type Person struct {
	first, last string
	favIceCream []string
}

func main() {
	cs := Person{
		first:       "Cyd",
		last:        "Salgado",
		favIceCream: []string{"Chocolate", "Cherry"},
	}

	as := Person{
		first:       "Alicia",
		last:        "Sutter",
		favIceCream: []string{"Vanilla", "Strawberry"},
	}

	// store the Person structs inside of a map
	ppl := map[string]Person{
		"salgado": cs,
		"sutter":  as,
	}

	// Access each value in the map
	// -- Print out the values
	// -- range over the slice - favIceCream
	fmt.Println(ppl["salgado"].first, ppl["salgado"].last)
	for _, v := range ppl["salgado"].favIceCream {
		fmt.Printf("\t%v\n", v)
	}

	fmt.Println(ppl["sutter"].first, ppl["sutter"].last)
	for _, v := range ppl["sutter"].favIceCream {
		fmt.Printf("\t%v\n", v)
	}
}