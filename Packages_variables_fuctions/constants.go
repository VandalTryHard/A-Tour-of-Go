package main

import "fmt"

const Pi = 3.14

func main() {
	const Word = "Golang"
	fmt.Println("Hello", Word)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
