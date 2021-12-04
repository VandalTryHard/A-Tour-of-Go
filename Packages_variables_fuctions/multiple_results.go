package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}

// package main

// import "fmt"

// func add(x, y string) (string, string) {
// 	return y, x
// }

// func main() {
// 	a, b := swap("golang", "Hello")
// 	fmt.Println(b, a)
// }
