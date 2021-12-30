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
	fmt.Println(lines)
}
