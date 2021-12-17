package main

import "fmt"

func main() {
	a := make([]int, 5)
	prinSlice("a", a)

	b := make([]int, 0, 5)
	prinSlice("b", b)

	c := b[:2]
	prinSlice("c", c)

	d := c[2:5]
	prinSlice("d", d)
}

func prinSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
