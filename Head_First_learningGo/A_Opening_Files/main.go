// Для работы с aardvark.txt
package main

import (
	"log"
	"os"
)

// Побитовый оператор "или" |
func main() {
	options := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	file, err := os.OpenFile("aardvark.txt", options, os.FileMode(0600))
	check(err)
	_, err = file.Write([]byte("Побитовый оператор 'или' |\n"))
	check(err)
	err = file.Close()
	check(err)

}

// Добаваление в конец файла__________________________________________________________________
// func main() {
// 	file, err := os.OpenFile("aardvark.txt", os.O_APPEND, os.FileMode(0600))
// 	check(err)
// 	_, err = file.Write([]byte("this os.O_APPEND\n")) // добавляет в конец
// 	check(err)
// 	err = file.Close()
// 	check(err)
// }

// Запись_____________________________________________________________________________________
// func main() {
// 	file, err := os.OpenFile("aardvark.txt", os.O_WRONLY, os.FileMode(0600))
// 	check(err)
// 	_, err = file.Write([]byte("os.O_WRONLY\n")) // Запись в файл
// 	check(err)
// 	err = file.Close()
// 	check(err)
// }

// Чтение_____________________________________________________________________________________
// func main() {
// 	file, err := os.OpenFile("aardvark.txt", os.O_RDONLY, os.FileMode(0600))
// 	check(err)
// 	defer file.Close()
// 	scanner := bufio.NewScanner(file)
// 	for scanner.Scan() {
// 		fmt.Println(scanner.Text())
// 	}
// 	check(scanner.Err())
// }

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
