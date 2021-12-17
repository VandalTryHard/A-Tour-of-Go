package main

import "fmt"

func main() {
	fmt.Println("couting")
	defer fmt.Println("done")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
}
