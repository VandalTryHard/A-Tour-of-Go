/*
По данному числу n закончите фразу "На лугу пасется..." одним из возможных продолжений: "n коров",
"n корова", "n коровы", правильно склоняя слово "корова".
Входные данные
Дано число n (0<n<100).
Выходные данные
Программа должна вывести введенное число n и одно из слов (на латинице): korov, korova или korovy,
например, 1 korova, 2 korovy, 5 korov. Между числом и словом должен стоять ровно один пробел.
Sample Input:
10

Sample Output:
10 korov
*/
package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	switch {
	case a == 1:
		fmt.Printf("%v korova", a)
	case a > 1 && a <= 4:
		fmt.Printf("%v korovy", a)
	case a >= 5 && a <= 20:
		fmt.Printf("%v korov", a)
	case a >= 21 && a%10 == 1:
		fmt.Printf("%v korova", a)
	case a >= 21 && (a%10 == 2 || a%10 == 3 || a%10 == 4):
		fmt.Printf("%v korovy", a)
	default:
		fmt.Printf("%v korov", a)
	}
}
