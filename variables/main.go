package main

import (
	"fmt"
	"math"
)

func main() {

	/* Variable declaration in a single statement
	var (
		name   = "naveen"
		age    = 29
		height int
	)
	fmt.Println("my name is", name, "age is", age, "and height is", height)
	*/

	/* Short hand declaration
	count := 10
	fmt.Println("Count =", count)
	*/

	/* Short hand multi variables
	name, age := "naveen", 29
	fmt.Println("my name is", name, "age is", age)
	*/

	/* Short hand assignment with values computed during run time*/
	a, b := 145.8, 543.8
	c := math.Min(a, b)
	fmt.Println("Minimum value is", c)

}
