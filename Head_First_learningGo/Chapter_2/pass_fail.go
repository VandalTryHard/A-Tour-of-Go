// pass_fail сообщает, сдал ли пользователь экзамен

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Print("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	if input >= 60 {
		status := "passing"
		fmt.Println(status)
	} else {
		status := "failing"
		fmt.Println(status)
	}
}
