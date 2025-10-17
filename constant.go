package main

import (
	"fmt"
	"math"
)

const s string = "constant"

// main demonstrates the use of constants in Go, including untyped numeric constants,
// arithmetic with constants, and type conversions. It prints the value of an undefined
// variable 's' (which will cause a compilation error if 's' is not defined elsewhere),
// calculates a large floating-point value 'd' using scientific notation and division,
// prints 'd' as a float and as an int64, and finally prints the sine of the constant 'n'
// using the math.Sin function.
func main() {
	fmt.Println(s)

	const n = 500000000

	const d = 3e20 / n
	fmt.Println(d)
	
	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}
