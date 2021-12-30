package main

import "fmt"

func main() {
	var notes []string = []string{"do", "re", "mi", "fa", "so", "la", "ti"}
	fmt.Println(len(notes))
	var primes []int = []int{2, 3, 5, 7, 11}
	fmt.Println(len(primes))

	//for
	fmt.Print("\n")
	for i := 0; i < len(notes); i++ {
		fmt.Println(i, notes[i])
	}

	// Range
	fmt.Print("\n")
	notesrange := []string{"do", "re", "mi", "fa", "so", "la", "ti"}
	for index, note := range notesrange {
		fmt.Println(index, note)
	}
}
