package main

import "fmt"

func main() {
	var a int
	var b int
	for a = 1; a <= 3; a++ {
		for b = 1; b <= 3; b++ {
			if a == 2 && b == 2 {
				break
			}
			fmt.Print(a, " ", b, "\n")
		}
	}
}
