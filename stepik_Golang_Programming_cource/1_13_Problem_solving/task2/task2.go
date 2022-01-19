/*
Дано трехзначное число. Переверните его, а затем выведите.
Формат входных данных
На вход дается трехзначное число, не оканчивающееся на ноль.
Формат выходных данных
Выведите перевернутое число.

Sample Input:
843

Sample Output:
348
*/
package main

import "fmt"

func main() {
	var a string
	fmt.Scan(&a)
	fmt.Printf("%v%v%v", string(a[2]), string(a[1]), string(a[0]))
}
