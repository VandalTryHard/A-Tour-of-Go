package main

import (
	"fmt"
	"time"
)

func a() {
	for i := 0; i < 100; i++ {
		fmt.Print("a")
	}
}
func b() {
	for i := 0; i < 100; i++ {
		fmt.Print("b")
	}
}
func main() {
	go a() //Конкурентно выполняется первой
	go b() //Конкурентно выполняется второй
	time.Sleep(time.Second)
	fmt.Println(" end main()")
}

// Решение задачи стр 422..................................................................................
// func main() {
// 	go repeat("x")
// 	go repeat("y")
// 	time.Sleep(time.Second)
// }

// func repeat(s string) {
// 	for i := 0; i <= 25; i++ {
// 		fmt.Print(s)
// 	}
// }

// использование канала chan...............................................................................
// func greeting(myChannel chan string) {
// 	myChannel <- "Hello myChannel"
// }

// func main() {
// 	myChannel := make(chan string)
// 	go greeting(myChannel)
// 	fmt.Println(<-myChannel)
// }

// Каналы и синхронизация горутин page 426................................................................
// func main() {
// 	channel1 := make(chan string)
// 	channel2 := make(chan string)
// 	go abc(channel1)
// 	go def(channel2)
// 	fmt.Print(<-channel1)
// 	fmt.Print(<-channel2)
// 	fmt.Print(<-channel1)
// 	fmt.Print(<-channel2)
// 	fmt.Print(<-channel1)
// 	fmt.Print(<-channel2)
// 	fmt.Println()
// }

// func abc(channel chan string) {
// 	channel <- "a"
// 	channel <- "b"
// 	channel <- "c"
// }

// func def(channel chan string) {
// 	channel <- "d"
// 	channel <- "e"
// 	channel <- "f"
// }

//Соблюдение синхронизации в горутине________________________________________________________________________
// func main() {
// 	myChannel := make(chan string)
// 	go send(myChannel)
// 	reportNap("receiving goroutine", 5)
// 	fmt.Println(<-myChannel)
// 	fmt.Println(<-myChannel)
// }

// func send(myChannel chan string) {
// 	reportNap("sending goroutine", 2)
// 	fmt.Println("***sending value***")
// 	myChannel <- "a" //Блокируется по этой отправке, пока «main» остается в ожидании.
// 	fmt.Println("***sending value***")
// 	myChannel <- "b"
// }

// func reportNap(name string, delay int) { //name-Имя прио-становленной горутины. delay-Время приостановки.
// 	for i := 0; i <= delay; i++ {
// 		fmt.Println(name, "sleeping")
// 		time.Sleep(1 * time.Second)
// 	}
// 	fmt.Println(name, "wakes up!")
// }

/* Заполните пропуски в коде, чтобы следующая программа использовала значения, ________________________________
полученные из двух каналов, и выводила показанный результат. Page 429
1
3
2
4
*/
// func main() {
// 	channelA := make(chan int)
// 	channelB := make(chan int)
// 	go odd(channelA)
// 	go even(channelB)
// 	fmt.Println(<-channelA)
// 	fmt.Println(<-channelA)
// 	fmt.Println(<-channelB)
// 	fmt.Println(<-channelB)
// }

// func odd(channel chan int) {
// 	channel <- 1
// 	channel <- 3
// }

// func even(channel chan int) {
// 	channel <- 2
// 	channel <- 4
// }
