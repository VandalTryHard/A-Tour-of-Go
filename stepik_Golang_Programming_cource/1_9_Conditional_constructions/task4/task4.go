/*
Дано неотрицательное целое число. Найдите и выведите первую цифру числа.
Формат входных данных
На вход дается натуральное число, не превосходящее 10000.
Формат выходных данных
Выведите одно целое число - первую цифру заданного числа.

Sample Input:
1234
Sample Output:
1
*/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a int = 1234
	// fmt.Scan(&a)
	str_a := strconv.Itoa(a)
	fmt.Println(string(str_a[0]))
}
