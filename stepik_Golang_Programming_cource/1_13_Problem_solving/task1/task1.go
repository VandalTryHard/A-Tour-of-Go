/*
Дано трехзначное число. Найдите сумму его цифр.
Формат входных данных
На вход дается трехзначное число.
Формат выходных данных
Выведите одно целое число - сумму цифр введенного числа.

Sample Input:
745

Sample Output:
16
*/
package main

import (
	"fmt"
)

func main() {
	var a, sum int
	fmt.Scan(&a)
	f := a / 100
	s := a / 10 % 10
	t := a % 10
	sum = f + s + t
	fmt.Println(sum)
}
