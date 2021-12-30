// Пакет datafile прдназначен для чтения данных из файлов.
package datafile

import (
	"bufio"
	"os"
	"strconv"
)

// GetFloats читает значение float64 из каждой строки  файла
func GetFloats(fileName string) ([3]float64, error) {
	var numbers [3]float64
	file, err := os.Open(fileName) // Открытие файла
	if err != nil {
		return numbers, err
	}
	i := 0                            // Переменная для хранения индексаЮ по которому должно выполняться присваивание
	scanner := bufio.NewScanner(file) // сканирование файла
	for scanner.Scan() {
		numbers[i], err = strconv.ParseFloat(scanner.Text(), 64) //Строка, прочитанная из файла, преобразуется в float64
		if err != nil {
			return numbers, err
		}
		i++
	}
	err = file.Close() // закрытие файла
	if err != nil {
		return numbers, err
	}
	if scanner.Err() != nil { //если при сканирование произошла ошибка
		return numbers, scanner.Err()
	}
	return numbers, nil
}
