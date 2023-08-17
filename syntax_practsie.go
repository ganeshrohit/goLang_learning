package main

import (
	"fmt"
	"math"
)

func add(a float64, b float64) float64 {
	return a + b
}

func main() {
	fmt.Println("The square of number 5 is : ", math.Pow(5, 2))

	var num1 float64 = 2.5
	var num2 float64 = 7.5

	/* Different ways to initialize the variables
	var num1, num2 float64  = 2.5, 7.5

	dynamic intialization
	num1, num2 := 2.5, 7.5
	*/

	fmt.Println("The sum of num1 and num2 is : ", add(num1, num2))

	x := num1

	fmt.Println("The value of x is: ", x)

	// Pointers
	a := 25
	ptra := &a

	fmt.Println(*ptra)
	*ptra = 30
	fmt.Println(*ptra, a)

	//strigs

	s1 := "hello"
	s2 := "there"

	fmt.Println(s1, s2)
}
