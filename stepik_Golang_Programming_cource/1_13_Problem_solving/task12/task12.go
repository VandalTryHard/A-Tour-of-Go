/*
По данному числу N распечатайте все целые значения степени двойки, не превосходящие N, в порядке
возрастания.
Входные данные
Вводится натуральное число.
Выходные данные
Выведите ответ на задачу.

Sample Input:
50

Sample Output:
1 2 4 8 16 32
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	var N int
	fmt.Scan(&N)

	p := float64(0)
	for {
		if math.Pow(2, p) <= float64(N) {
			fmt.Printf("%v ", math.Pow(2, p))
			p++
		} else {
			break
		}
	}
}
