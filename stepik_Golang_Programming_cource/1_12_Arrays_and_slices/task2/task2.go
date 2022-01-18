/*
Напишите программу, принимающая на вход число N(N≥4), а затем N целых чисел-элементов среза.
На вывод нужно подать 4-ый (3 по индексу) элемент данного среза.

Sample Input:
5
41 -231 24 49 6

Sample Output:
49
*/
package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	mySlice := make([]int, N)
	for i := 0; i < N; i++ {
		var num int
		fmt.Scan(&num)
		mySlice[i] = num
	}
	fmt.Println(mySlice[3])
}
