package main


import "fmt"

type person struct {
	first string
	license int
	sayings []string
}

func main() {
	p1 := person{
		first: "James",
		license: 007,
		sayings: []string{"Shaken, not stirred", "Bond, James Bond",},
	}
	fmt.Println(p1)

	p2 := person{
		first: "Jenny",
		license: 8,
		sayings: []string{"whatever she says", "Bond, Oh hes my man",},
	}
	fmt.Println(p2)

	xp := []person{p1, p2}
	fmt.Println(xp)

	mp := map[string]person{"Mr":p1, "Ms":p2,}
	fmt.Print(mp)
	

	
}