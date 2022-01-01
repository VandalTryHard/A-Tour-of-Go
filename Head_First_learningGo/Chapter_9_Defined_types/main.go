package main

import (
	"fmt"
)

type Liters float64
type Milliliters float64
type Gallons float64

// type MyType string

func (l Liters) ToGallons() Gallons {
	return Gallons(l * 0.264)
}

func (m Milliliters) ToGallons() Gallons {
	return Gallons(m * 0.000264)
}

func (g Gallons) ToLiters() Liters {
	return Liters(g * 4.546)
}

func main() {
	soda := Liters(2)
	fmt.Printf("%0.3f liters equals %0.3f gallons\n", soda, soda.ToGallons())

	water := Milliliters(500)
	fmt.Printf("%0.3f milliliters equals %0.3f gallons\n", water, water.ToGallons())

	carFuel := Gallons(1.2)
	busFuel := Liters(2.5)
	// carFuel += ToGallons(Liters(40))
	// busFuel += ToLiters(Gallons(30.0))
	fmt.Printf("%0.3f gallons equals %0.3f liters\n", carFuel, carFuel.ToLiters())
	fmt.Printf("%0.3f liters equals %0.3f gallons\n", busFuel, busFuel.ToGallons())
	// fmt.Printf("Car fuel: %0.1f gallons\n", carFuel)
	// fmt.Printf("Bus fuel: %0.1f liters\n", busFuel)

	// value := MyType("MyType value")
	// value.MethodWithParameters(4, true)
}

// func (m MyType) MethodWithParameters(number int, flag bool) {
// 	fmt.Println(m)
// 	fmt.Println(number)
// 	fmt.Println(flag)
// }

// // Imports, type declarations omitted
// func ToGallons(l Liters) Gallons {
// 	return Gallons(l * 0.264)
// }
// func ToLiters(g Gallons) Liters {
// 	return Liters(g * 3.785)
// }
