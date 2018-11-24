/*
Create a for loop using this syntax
	for {}
Have it print out th eyars you have been alive.
*/
package main

import "fmt"

func main() {
	birthYear := 1971
	curYear := 2018
	i := birthYear

	for {
		if i <= curYear {
			fmt.Println(i)
			i++
		} else {
			break
		}
	}
	fmt.Println("Still kicking!")
}