package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	fmt.Printf("Sqrt approximation of %v:\n", x)
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Printf("Iteration %v, value = %v\n", i, z)
	}
	return z
}

func checkSqrtMath(x float64) float64 {
	fmt.Printf("Sqrt approximation of %v:\n", x)
	c := math.Sqrt(x)
	fmt.Printf("math.Sqrt(x) value = %v\n", c)
	return c
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(checkSqrtMath(2))
}
