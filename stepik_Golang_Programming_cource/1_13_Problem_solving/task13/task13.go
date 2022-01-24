/*
Номер числа Фибоначчи
Дано натуральное число A > 1. Определите, каким по счету числом Фибоначчи оно является, то есть выведите такое число n, что φn=A.
Если А не является числом Фибоначчи, выведите число -1.
Входные данные
Вводится натуральное число.
Выходные данные
Выведите ответ на задачу.

Sample Input:
8

Sample Output:
6
*/
package main

import "fmt"

func main() {
	var A int
	fmt.Scan(&A)
	nFib, notFib := 1, -1
	a, b := 1, 1
	for a != A {
		a, b = b, a+b
		nFib++
		if a == A {
			fmt.Println(nFib)
		} else if a > A {
			fmt.Println(notFib)
			break
		}
	}
}
