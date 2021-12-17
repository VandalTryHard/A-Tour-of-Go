package main

import "fmt"

// k := 3 dont work
var k int = 3 //work

func main() {
	var x, y int = 1, 2
	// k := 3 //work
	var c, python, java = false, "-my favorite language is python-", "java"
	fmt.Println(x, y, k, c, python, java)
}
