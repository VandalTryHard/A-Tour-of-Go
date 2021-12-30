// Программа должна объявлять целочисленную переменную myInt и переменную
// myIntPointer для хранения целочисленного указателя. Затем она должна присваивать
// значение myInt и указатель на myInt как значение myIntPointer. Наконец, программа
// должна выводить значение по указателю myIntPointer.
//https://www.youtube.com/watch?v=geDH2K9Hh8w
package main

import "fmt"

func main() {
	var myInt int
	var myIntPointer *int //выделяем памяит
	fmt.Println(myIntPointer)
	myInt = 42
	fmt.Println(myInt, &myInt)
	myIntPointer = &myInt // указываем myIntPointer значение myInt в выделенной памяти
	*myIntPointer = 5     // изменяем значение myInt по адресу
	fmt.Println(*myIntPointer, myIntPointer)
}
