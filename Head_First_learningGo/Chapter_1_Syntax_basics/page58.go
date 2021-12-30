//Преобразование типов

package main

import "fmt"

func main() {
	var price int = 100
	fmt.Println("Price is", price, "dollars.")

	var taxRate float64 = 0.08
	var tax float64 = float64(price) * taxRate
	fmt.Println("Tax is", tax, "dollars")

	var total float64 = tax + float64(price)
	fmt.Println("Total cost is", int(total), "dollars.")

	var availableFunds int = 120
	fmt.Println(availableFunds, "dollars available.")
	fmt.Println("Within budhet?", availableFunds > int(total))
}
