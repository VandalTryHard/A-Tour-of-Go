/*
Из натурального числа удалить заданную цифру.
Входные данные
Вводятся натуральное число и цифра, которую нужно удалить.
Выходные данные
Вывести число без заданных цифр.
Sample Input:
38012732
3

Sample Output:
801272
*/
package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	result, rank := 0, 1
	for a != 0 {
		reminder := a % 10
		if reminder != b {
			result += reminder * rank
			rank *= 10
		}
		a /= 10
	}
	fmt.Println(result)
}
