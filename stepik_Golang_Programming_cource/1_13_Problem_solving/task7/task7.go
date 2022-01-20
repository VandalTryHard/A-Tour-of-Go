/*
Количество нулей
По данным числам, определите количество чисел, которые равны нулю.
Входные данные
Вводится натуральное число N, а затем N чисел.
Выходные данные
Выведите количество чисел, которые равны нулю.

Sample Input:
5
1
8
100
0
12

Sample Output:
1
*/
package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	var zero int
	for i := 1; i <= N; i++ {
		var a int
		fmt.Scan(&a)
		if a == 0 {
			zero += 1
		}
	}
	fmt.Println(zero)
}
