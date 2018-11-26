/*
Create a func which returns a func
assign the returned func to a variable
call the returned func
*/
package main

import (
	"fmt"
)

func main() {

	sayHiHenry := greeting("Henry")
	sayHiSue := greeting("Sue")

	fmt.Println(sayHiHenry())
	fmt.Println(sayHiSue())
}

func greeting(name string) func() string {
	return func() string {
		return fmt.Sprintf("Hello, my name is %v", name)
	}
}