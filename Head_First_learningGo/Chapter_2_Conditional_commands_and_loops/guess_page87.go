// guess - игра, в которой игрок должен угадать случайное число. Page 89 - page 107
package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	second := time.Now().Unix()
	rand.Seed(second)
	target := rand.Intn(100) + 1
	fmt.Println("I've chosen a random number between 1 and 100.")
	fmt.Println("Can you guess it?")
	//fmt.Println(target)

	reader := bufio.NewReader(os.Stdin)

	success := false
	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("You have", 10-guesses, "guesses left.")
		fmt.Println("Make a quess: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		if guess < target {
			fmt.Println("Oops. Your guess was LOW!")
		} else if guess > target {
			fmt.Println("Oops. Your guess was HIGH!")
		} else {
			success = true
			fmt.Println("Good job! You guessed it!")
			break
		}
	}
	if !success {
		fmt.Println("Sorry. You didn't guess my number. It was:", target)
	}
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
