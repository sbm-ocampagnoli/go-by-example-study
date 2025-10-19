package main

import (
	"fmt"
)

func main() {

	// Array declaration and initialization by default is zero-valued
	var a [5]int
	fmt.Println("emp:", a)
	// Set Value array index 4 and get value at index 4
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])
	// Get Length of array
	fmt.Println("len:", len(a))
	//Use this syntax to declare and initialize an array in one line.
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	//You can also have the compiler count the number of elements for you with ...
	c := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println("dcl:", c)

	//If you specify the index with :, the elements in between will be zeroed.
	d := [5]int{100, 3: 400, 500}
	fmt.Println("dcl:", d)

	var twoD [2][3]int
	for i := range 2 {
		for j := range 3 {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	twoD = [2][3]int{
		{0, 1, 2},
		{0, 1, 2},
	}
	fmt.Println("2d: ", twoD)
}
