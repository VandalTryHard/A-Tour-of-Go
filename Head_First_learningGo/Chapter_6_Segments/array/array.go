package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("\n", primes())

	var dates [3]time.Time
	dates[0] = time.Unix(1257894000, 0)
	dates[1] = time.Unix(1447920000, 0)
	dates[2] = time.Unix(1508632200, 0)
	fmt.Println(dates[1])

	var notes [7]string = [7]string{"do", "re", "mi", "fa", "so", "la", "ti"}
	var primes [5]int = [5]int{2, 3, 5, 7, 11}
	fmt.Printf("%#v\n", notes)
	fmt.Printf("%#v\n", primes)

	index := 1
	fmt.Println(index, notes[index])
	index = 3
	fmt.Println(index, notes[index])

	fmt.Println("\nfor для notes page 189")
	for i := 0; i <= 2; i++ {
		fmt.Println(i, notes[i])
	}
}

func primes() []int {
	var primes [5]int
	primes[0] = 1
	primes[3] = 4
	return primes[:]
}
