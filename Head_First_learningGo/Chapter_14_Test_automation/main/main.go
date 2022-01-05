package main

import (
	"fmt"

	"../prose"
)

func main() {
	phrarses := []string{"my parents", "a rodeo clown"}
	fmt.Println("A photo of", prose.JoinWithCommas(phrarses))
	phrarses = []string{"my parents", "a rodeo clown", "a prize bull"}
	fmt.Println("A photo of", prose.JoinWithCommas(phrarses))
	phrarses = []string{"my parents"}
	fmt.Println("A photo of", prose.JoinWithCommas(phrarses))
}
