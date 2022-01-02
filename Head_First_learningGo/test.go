package main

import "fmt"

func main() {
	var price float64 = 100
	var percentYear float64 = 0.3
	var value float64 = 0
	for i := 0; i < 365; i++ {
		value += price * (percentYear / 365.0)
		fmt.Printf("Day: %d, value %.2f\n", i, price*(percentYear/365.0))
	}
	fmt.Printf("Year value: %.3f", value)
}
