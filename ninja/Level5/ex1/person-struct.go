/*
Create your own type “person” which will have an underlying type of “struct”
so that it can store the following data:
  - first name
  - last name
  - favorite ice cream flavors
Create two VALUES of TYPE person. Print out the values, ranging over the
elements in the slice which stores the favorite flavors.
*/
package main

import (
	"fmt"

	"github.com/duld/learningGolang/exercises/learning-golang/level_5/person"
)

//Person : a representation of a person
type Person struct {
	first, last string
	favIceCream []string
}

func main() {
	mp := person.LPerson{
		First:       "Sal",
		Last:        "Esparza",
		FavIceCream: []string{"a"},
	}

	fmt.Println(mp)

	jm := Person{
		first:       "Cyd",
		last:        "Salgado",
		favIceCream: []string{"Vanilla", "Rocky Road"},
	}

	kf := Person{
		first:       "Alicia",
		last:        "Sutter",
		favIceCream: []string{"Moose Tracks", "Strawberry"},
	}

	ppl := []Person{jm, kf}

	for _, v := range ppl {
		fmt.Println(v.first, v.last)
		for _, flavor := range v.favIceCream {
			fmt.Printf("\t%v\n", flavor)
		}
	}
}