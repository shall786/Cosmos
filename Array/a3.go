/*

Go Slice
In Go, slice is a dynamically-sized, segmented view of an underlying array.
This segment can be the entire array or a subset of an array.
We define the subset of an array by indicating the start and end index.
Slices provide a dynamic window onto the underlying array.

*/

package main

import (
	"fmt"
)

func main() {
	odd := [6]int{2, 4, 6, 8, 10, 12}
	var s []int = odd[1:4]
	fmt.Println(s)
}
