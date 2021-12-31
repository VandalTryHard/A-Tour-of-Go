// count подсчитывает количество вхождений каждой строки в файле.
package main

import (
	"fmt"
	"log"

	"../datafile"
)

func main() {
	lines, err := datafile.GetStrings("votes.txt")
	if err != nil {
		log.Fatal(err)
	}

	// с использованием map
	counts := make(map[string]int)
	for _, line := range lines {
		counts[line]++
	}

	for string, i := range counts {
		fmt.Printf("%s: %v\n", string, i)
	}

	// с использованием 2х переменных
	// var names []string
	// var counts []int
	// for _, line := range lines {
	// 	matched := false
	// 	for i, name := range names {
	// 		if name == line {
	// 			counts[i]++
	// 			matched = true
	// 		}
	// 	}

	// 	if matched == false {
	// 		names = append(names, line)
	// 		counts = append(counts, 1)
	// 	}
	// }

	// for i, name := range names {
	// 	fmt.Printf("%s: %d\n", name, counts[i])
	// }

	// fmt.Println(lines)
}
