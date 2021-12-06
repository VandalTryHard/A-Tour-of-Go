package main

import "fmt"

func main() {
	prime := [6]int{1, 2, 3, 4, 5, 6}

	var s []int = prime[1:4]
	var a []int = prime[1:6]
	fmt.Println(s, a)
	x := append(s, a...)
	fmt.Println(x)
}
