package main

import "fmt"

func main() {
	/* local variable definition */
	var x int = 10
	var y int = 20
	/* check the boolean condition */
	if x >= 10 {
		/* if condition is true then check the following */
		if y >= 10 {
			/* if condition is true then print the following */
			fmt.Printf("Inside nested If Statement \n")
		}
	}
	fmt.Printf("Value of x is : %d\n", x)
	fmt.Printf("Value of y is : %d\n", y)
}
