// Пакет datafile прдназначен для чтения данных из файлов.
package datafile

import (
	"bufio"
	"os"
	"strconv"
)

// GetFloats читает значение float64 из каждой строки  файла
func GetFloats(fileName string) ([]float64, error) {
	var numbers []float64
	file, err := os.Open(fileName) // Открытие файла
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file) // сканирование файла
	for scanner.Scan() {
		number, err := strconv.ParseFloat(scanner.Text(), 64) //Строка, прочитанная из файла, преобразуется в float64
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, number)
	}
	err = file.Close() // закрытие файла
	if err != nil {
		return nil, err
	}
	if scanner.Err() != nil { //если при сканирование произошла ошибка
		return nil, scanner.Err()
	}
	return numbers, nil
}
