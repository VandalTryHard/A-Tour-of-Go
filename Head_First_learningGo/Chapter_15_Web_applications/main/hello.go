// Простые ответы браузеру
package main

import (
	"log"
	"net/http"
)

func write(writer http.ResponseWriter, message string) {
	_, err := writer.Write([]byte(message))
	if err != nil {
		log.Fatal(err)
	}
}

func englishHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Hello, web!")
}

func frenchHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Salut web!")
}

func hindiHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Namaste, web!")
}

func main() {
	http.HandleFunc("/hello", englishHandler)
	http.HandleFunc("/salut", frenchHandler)
	http.HandleFunc("/namaste", hindiHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}

// Task page 473______________________________________________________________________________________
// import "fmt"

// func callFunction(passedFunction func()) {
// 	passedFunction()
// }
// func callTwice(passedFunction func()) {
// 	passedFunction()
// 	passedFunction()
// }
// func callWithArguments(passedFunction func(string, bool)) {
// 	passedFunction("This sentence is", false)
// }
// func printReturnValue(passedFunction func() string) {
// 	fmt.Println(passedFunction())
// }
// func functionA() {
// 	fmt.Println("function called")
// }
// func functionB() string {
// 	fmt.Println("function called")
// 	return "Returning from function"
// }
// func functionC(a string, b bool) {
// 	fmt.Println("function called")
// 	fmt.Println(a, b)
// }
// func main() {
// 	callFunction(functionA)
// 	callTwice(functionA)
// 	callWithArguments(functionC)
// 	printReturnValue(functionB)
// }

// Рефакторинг Page 467_______________________________________________________________________________
// func viewHandler(writer http.ResponseWriter, request *http.Request) {
// 	message := []byte("Hello, web!")
// 	_, err := writer.Write(message)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// func main() {
// 	http.HandleFunc("/hello", viewHandler)
// 	err := http.ListenAndServe("localhost:8080", nil)
// 	log.Fatal(err)
// }
