package main

import "fmt"

func main() {

	// simple condition
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// A classic initial/condition/after for loop
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}
	// other way
	for i := range 3 {
		fmt.Println("range", i)
	}
	// withour condition using break or return
	for {
		fmt.Println("loop")
		break
	}
	// for continue to the next interation of loop
	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
