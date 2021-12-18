// Игра отгадай число Page 89
package main

import "math/rand"

func main() {
	target := rand.Intn(100) + 1
}

// package main

// import (
// 	"bufio"
// 	"crypto/rand"
// 	"fmt"
// 	"log"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// func maing() {
// 	var try int
// 	num := rand.Intn(100) + 1
// 	for try = 10; try > 0; try-- {
// 		fmt.Println("У тебя есть", try, "попыток")
// 		fmt.Print("Enter number: ")
// 		reader := bufio.NewReader(os.Stdin)
// 		input, err := reader.ReadString('n')
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		input = strings.TrimSpace(input)
// 		grade, err := strconv.ParseFloat(input, 64)
// 		var status string
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		if grade < num {
// 			status = "Oops. Your guess was LOW"
// 		} else if grade > num {
// 			status = "Oops. Your guess was HIGH"
// 		} else {
// 			break
// 		}
// 		fmt.Println(status)
// 	}
// }
